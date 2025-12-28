package mp3

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/camggould/aqa/validation"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	ffprobe "gopkg.in/vansante/go-ffprobe.v2"
)

type Mp3 struct {
	filePath string
	samples []float64
	sampleRate int
}

const HANN_RMS = 0.6123724356957945

func (*Mp3) New(filePath string) (*Mp3, error) {

	err := validation.IsMP3(filePath)
	if err != nil {
		return nil, err
	}

	sampleRate, err := detectSampleRate(filePath)

	if err != nil {
		return nil, fmt.Errorf("No sample rate detected for provided mp3 file %s. %w", filePath, err)
	}

	samples, err := decodeToPCM(filePath, sampleRate)

	if err != nil {
		return nil, fmt.Errorf("Failed to collect samples for file %s. %w", filePath, err)
	}

	sampleRateAsInt, err := parseSampleRate(sampleRate)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse sample rate [%s]. %w", sampleRate, err)
	}

	mp3 := Mp3 {
		filePath: filePath,
		samples: samples,
		sampleRate: sampleRateAsInt,
	}

	return &mp3, nil

}

func (mp3 *Mp3) GetRmsFloor() float64 {
	frameSize := int(0.5 * float64(mp3.sampleRate))
	hopSize := int(0.01 * float64(mp3.sampleRate)) // e.g., 10 ms hop

	minRMS := math.Inf(1)
	for i := 0; i+frameSize <= len(mp3.samples); i += hopSize {
		windowed := applyHannWindow(mp3.samples[i : i+frameSize])
		rms := rmsFrame(windowed)
		if rms < minRMS {
			minRMS = rms
		}
	}

    return rmsToDBFS(minRMS)
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
    for _, v := range samples {
        sum += v
    }
    mean := sum / float64(len(samples))

    var sq float64
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
		})

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

	// Defensive validation — realistic audio ranges
	if value < 8000 || value > 384000 {
		return 0, fmt.Errorf("unreasonable sample rate: %d", value)
	}

	return value, nil
}

func rmsToDBFS(rms float64) float64 {
	if rms <= 0 {
		return math.Inf(-1)
	}
	return 20 * math.Log10(rms)
}
