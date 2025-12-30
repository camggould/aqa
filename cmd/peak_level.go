package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/utils"

	"github.com/spf13/cobra"
)

type PeakResponse struct {
	File string `json:"file"`
	PeakLevel string `json:"peakLevel"`
}

var peakLevelCmd = &cobra.Command{
	Use: "peak",
	Short: "Get the peak level of an audio file.",
	Long: "Get the peak level of a provided mp3 file. Must be a valid audio file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(HandleAudioAnalysis(cmd, args, runPeakLevelCommand))
		return
	},
}

func runPeakLevelCommand(cmd *cobra.Command, args []string, audioFile string) string {
	peakLevel, err := GetPeakLevel(audioFile)

	if err != nil {
		return fmt.Sprintf("Peak level could not be calculated due to error: %s\n", err)
	}

	responseData := PeakResponse{
		File: audioFile,
		PeakLevel: utils.PrintDb(peakLevel),
	}

	return utils.FormattedJsonOutput(responseData)
}

func GetPeakLevel(filePath string) (float64, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetPeakDBFS(), nil
}