package cmd

import (
	"fmt"

	"github.com/camggould/aqa/mp3"

	"github.com/spf13/cobra"
)

var rmsFloorCmd = &cobra.Command{
	Use: "rmsFloor",
	Short: "Get the RMS floor of an mp3 file.",
	Long: "Get the RMS floor of a provided mp3 file. Must be a valid MP3. Calculated based on the quietest 0.5s segment of audio.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(HandleMp3Analysis(cmd, args, runRmsFloorCommand))
		return
	},
}

func runRmsFloorCommand(cmd *cobra.Command, args []string, mp3File string) string {
	rmsFloor, err := GetRmsFloor(mp3File)

	if err != nil {
		return fmt.Sprintf("RMS floor could not be calculated due to error: %s\n", err)
	}

	return fmt.Sprintf("RMS floor for file %s: %fdB\n", mp3File, rmsFloor)
}

func init() {
	rootCmd.AddCommand(rmsFloorCmd)
}

func GetRmsFloor(filePath string) (float64, error) {
	var audio *mp3.Mp3

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetRmsFloor(), nil

}