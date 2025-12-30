package audio

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

type AudioMetadata struct {
	Channels int
	SampleRate int
}

func (*AudioMetadata) New(path string) (audioMetadata *AudioMetadata, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil { return }

	sampleRate, err := sampleRateToInt(getSampleRate(data.Streams))
	if err != nil { return }

	channelCount, err := getChannelCount(data.Streams)
	if err != nil { return }

	audioMetadata = &AudioMetadata{
		Channels: channelCount,
		SampleRate: sampleRate,
	}

	return
}

func getSampleRate(streams []*ffprobe.Stream) (sampleRate string, err error) {
	for _, stream := range streams {
		if stream.CodecType == "audio" {
			if stream.SampleRate == "" {
				return "", fmt.Errorf("missing sample_rate")
			}
			return stream.SampleRate, nil
		}
	}

	return "", fmt.Errorf("No sample rate found.")
}

func getChannelCount(streams []*ffprobe.Stream) (channelCount int, err error) {
	for _, stream := range streams {
		if stream.CodecType == "audio" {
			return stream.Channels, nil
		}
	}

	return 0, fmt.Errorf("No channel count found.")
}

func sampleRateToInt(sr string, inputErr error) (int, error) {
	if inputErr != nil { return 0, inputErr }

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