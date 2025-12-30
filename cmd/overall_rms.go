package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"
	"github.com/spf13/cobra"
)

var rmsCmd = &cobra.Command{
	Use: "rms",
	Short: "Get the RMS of an audio file.",
	Long: "Get the RMS of a provided audio file. Must be a valid audio file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(HandleAudioAnalysis(cmd, args, runRmsCommand))
		return
	},
}

func runRmsCommand(cmd *cobra.Command, args []string, audioFile string) string {
	rms, err := GetOverallRMS(audioFile)

	if err != nil {
		return fmt.Sprintf("RMS could not be calculated due to error: %s\n", err)
	}

	return fmt.Sprintf("RMS for file %s: %fdB\n", audioFile, rms)
}

func init() {
	rootCmd.AddCommand(rmsCmd)
}

func GetOverallRMS(filePath string) (float64, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetOverallRMS(), nil
}