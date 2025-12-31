package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/utils"

	"github.com/spf13/cobra"
)

type SampleRateResponse struct {
	File string `json:"file"`
	SampleRate string `json:"sampleRate"`
}

var sampleRateCmd = &cobra.Command{
	Use: "sampleRate",
	Short: "Get the sample rate of an audio file.",
	Long: "Get the sample rate of a provided mp3 file. Must be a valid audio file.",
	Run: RunGenerator(runSampleRateCommand),
}

func runSampleRateCommand(cmd *cobra.Command, args []string, audioFile string) string {
	sampleRate, err := GetSampleRate(audioFile)

	if err != nil {
		return fmt.Sprintf("Sample rate could not be calculated due to error: %s\n", err)
	}

	responseData := SampleRateResponse{
		File: audioFile,
		SampleRate: utils.PrintSampleRate(sampleRate),
	}

	return utils.FormattedJsonOutput(responseData)
}

func GetSampleRate(filePath string) (int, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetSampleRate(), nil
}