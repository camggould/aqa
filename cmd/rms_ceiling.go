package cmd

import (
	"fmt"

	"github.com/camggould/aqa/mp3"

	"github.com/spf13/cobra"
)

var rmsCeilingCmd = &cobra.Command{
	Use: "rmsCeiling",
	Short: "Get the RMS ceiling of an mp3 file.",
	Long: "Get the RMS ceiling of a provided mp3 file. Must be a valid MP3. Calculated based on the loudest 0.5s segment of audio.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rmsFloor, err := GetRmsCeiling(args[0])

		if err != nil {
			fmt.Printf("RMS ceiling could not be calculated due to error: %s\n", err)
			return
		}

		fmt.Printf("RMS ceiling for file %s: %fdB\n", args[0], rmsFloor)
	},
}

func init() {
	rootCmd.AddCommand(rmsCeilingCmd)
}

func GetRmsCeiling(filePath string) (float64, error) {
	var audio *mp3.Mp3

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetRMSCeiling(), nil

}