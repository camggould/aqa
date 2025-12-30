package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"

	"github.com/spf13/cobra"
)

var rmsFloorCmd = &cobra.Command{
	Use: "rmsFloor",
	Short: "Get the RMS floor of an audio file.",
	Long: "Get the RMS floor of a provided audio file. Must be a valid audio file. Calculated based on the quietest 0.5s segment of audio.",
	Run: RunGenerator(runRmsFloorCommand),
}

func runRmsFloorCommand(cmd *cobra.Command, args []string, audioFile string) string {
	rmsFloor, err := GetRmsFloor(audioFile)

	if err != nil {
		return fmt.Sprintf("RMS floor could not be calculated due to error: %s\n", err)
	}

	return fmt.Sprintf("RMS floor for file %s: %fdB\n", audioFile, rmsFloor)
}

func GetRmsFloor(filePath string) (float64, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetRmsFloor(), nil

}