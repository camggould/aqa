package audio

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	ffprobe "gopkg.in/vansante/go-ffprobe.v2"
)

type AudioFile struct {
	filePath string
	samples []float64
	sampleRate int
	rms float64
	peak float64
}

const HANN_RMS = 0.6123724356957945

func (*AudioFile) New(filePath string) (*AudioFile, error) {
	sampleRate, err := detectSampleRate(filePath)

	if err != nil {
		return nil, fmt.Errorf("No sample rate detected for provided audio file %s. %w", filePath, err)
	}

	samples, err := decodeToPCM(filePath, sampleRate)

	if err != nil {
		return nil, fmt.Errorf("Failed to collect samples for file %s. %w", filePath, err)
	}

	sampleRateAsInt, err := parseSampleRate(sampleRate)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse sample rate [%s]. %w", sampleRate, err)
	}

	rms, peak, err := parseAudioLevels(filePath)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse rms or peak for file [%s]. %w", filePath, err)
	}

	audioFile := AudioFile {
		filePath: filePath,
		samples: samples,
		sampleRate: sampleRateAsInt,
		rms: rms,
		peak: peak,
	}

	return &audioFile, nil
}

/** Calculates overall RMS
 * Overall RMS is calculated by leveraging ffmpeg's volumedetect to find the mean volume.
*/
func (audioFile *AudioFile) GetOverallRMS() float64 {
	return audioFile.rms
}

/** Calculates RMS floor
 * RMS floor is calculated by sliding across 0.5s windows of the audio (0.1s hops)
 * and calculating the RMS of that window. The minimum RMS window is tracked and converted
 * to decibles at the end. Each window that is evaluated includes DC removal and Hann smoothing.
*/
func (audioFile *AudioFile) GetRmsFloor() float64 {
	frameSize := int(0.5 * float64(audioFile.sampleRate))
	hopSize := int(0.01 * float64(audioFile.sampleRate))

	minRMS := math.Inf(1)
	for i := 0; i+frameSize <= len(audioFile.samples); i += hopSize {
		windowed := applyHannWindow(audioFile.samples[i : i+frameSize])
		rms := rmsFrame(windowed)
		if rms < minRMS {
			minRMS = rms
		}
	}

	return rmsToDBFS(minRMS)
}

/** Calculates RMS ceiling
 * RMS ceiling is calculated by sliding across 0.5s windows of the audio (0.1s hops)
 * and calculating the RMS of that window. The maximum RMS window is tracked and converted
 * to decibles at the end. Each window that is evaluated includes DC removal and Hann smoothing.
*/
func (audioFile *AudioFile) GetRMSCeiling() float64 {
	frameSize := int(0.5 * float64(audioFile.sampleRate))
	hopSize := int(0.01 * float64(audioFile.sampleRate))

	maxRMS := math.Inf(-1)

	for i := 0; i+frameSize <= len(audioFile.samples); i += hopSize {
		frame := make([]float64, frameSize)
		copy(frame, audioFile.samples[i:i+frameSize])

		windowed := applyHannWindow(frame)
		rms := rmsFrame(windowed)

		if rms > maxRMS {
			maxRMS = rms
		}
	}

	return rmsToDBFS(maxRMS)
}

/* Returns peak level of audio.
 * This leverages ffmpeg volumedetect's max_volume to retrieve the peak volume.
 */
func (audioFile *AudioFile) GetPeakDBFS() float64 {
	return audioFile.peak
}

func applyHannWindow(samples []float64) []float64 {
	n := len(samples)
	out := make([]float64, n)

	for i := 0; i < n; i++ {
		w := 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(n-1)))
		out[i] = samples[i] * (w / HANN_RMS)
	}

	return out
}

func rmsFrame(samples []float64) float64 {
	var sum float64
	var sq float64
	for _, v := range samples {
		sum += v
	}
	mean := sum / float64(len(samples))

	for _, v := range samples {
		x := v - mean
		sq += x * x
	}
	return math.Sqrt(sq / float64(len(samples)))
}

func detectSampleRate(path string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		return "", err
	}

	for _, stream := range data.Streams {
		if stream.CodecType == "audio" {
			if stream.SampleRate == "" {
				return "", fmt.Errorf("missing sample_rate")
			}
			return stream.SampleRate, nil
		}
	}

	return "", fmt.Errorf("no audio stream found")
}

func decodeToPCM(filePath string, sampleRate string) ([]float64, error) {
	cmd := ffmpeg.Input(filePath).
		Output("pipe:", ffmpeg.KwArgs{
			"format": "f32le",
			"nostats": "",
			"loglevel": "error",
			"f": "f32le",
			"ac": 1,
			"ar": sampleRate,
		}).
		Silent(true)

	var out bytes.Buffer

	err := cmd.WithOutput(&out, nil).Run()
	if err != nil {
		return nil, err
	}

	raw := out.Bytes()
	samples := make([]float64, len(raw)/4)

	for i := 0; i < len(samples); i++ {
		bits := binary.LittleEndian.Uint32(raw[i*4:])
		samples[i] = float64(math.Float32frombits(bits))
	}

	return samples, nil
}

func parseSampleRate(sr string) (int, error) {
	if sr == "" {
		return 0, fmt.Errorf("empty sample rate")
	}

	value, err := strconv.Atoi(sr)
	if err != nil {
		return 0, fmt.Errorf("invalid sample rate %q: %w", sr, err)
	}

	if value < 8000 || value > 384000 {
		return 0, fmt.Errorf("unreasonable sample rate: %d", value)
	}

	return value, nil
}

func parseAudioLevels(filePath string) (meanDB float64, maxDB float64, err error) {
	var stderr bytes.Buffer

	err = ffmpeg.Input(filePath).
		Output("pipe:", ffmpeg.KwArgs{
			"af": "volumedetect",
			"f":  "null",
		}).
		WithOutput(nil, &stderr).
		Silent(true).
		Run()
	if err != nil {
		return
	}

	log := stderr.String()

	reMean := regexp.MustCompile(`mean_volume: ([\-\d\.]+) dB`)
	reMax  := regexp.MustCompile(`max_volume: ([\-\d\.]+) dB`)

	meanDB, _ = strconv.ParseFloat(reMean.FindStringSubmatch(log)[1], 64)
	maxDB, _  = strconv.ParseFloat(reMax.FindStringSubmatch(log)[1], 64)

	return
}

func rmsToDBFS(rms float64) float64 {
	if rms <= 0 {
		return math.Inf(-1)
	}
	return 20 * math.Log10(rms)
}
