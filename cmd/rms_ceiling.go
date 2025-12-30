package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"

	"github.com/spf13/cobra"
)

var rmsCeilingCmd = &cobra.Command{
	Use: "rmsCeiling",
	Short: "Get the RMS ceiling of an audio file.",
	Long: "Get the RMS ceiling of a provided audio file. Must be a valid audio file. Calculated based on the loudest 0.5s segment of audio.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(HandleAudioAnalysis(cmd, args, runRmsCeilingCommand))
		return
	},
}

func runRmsCeilingCommand(cmd *cobra.Command, args []string, audioFile string) string {
	rmsFloor, err := GetRmsCeiling(audioFile)

	if err != nil {
		return fmt.Sprintf("RMS ceiling could not be calculated due to error: %s\n", err)
	}

	return fmt.Sprintf("RMS ceiling for file %s: %fdB\n", audioFile, rmsFloor)
}

func GetRmsCeiling(filePath string) (float64, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetRMSCeiling(), nil

}