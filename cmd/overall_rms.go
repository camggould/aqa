package cmd

import (
	"fmt"

	"github.com/camggould/aqa/mp3"

	"github.com/spf13/cobra"
)

var rmsCmd = &cobra.Command{
	Use: "rms",
	Short: "Get the RMS of an mp3 file.",
	Long: "Get the RMS of a provided mp3 file. Must be a valid MP3.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rms, err := GetOverallRMS(args[0])

		if err != nil {
			fmt.Printf("RMS could not be calculated due to error: %s\n", err)
			return
		}

		fmt.Printf("RMS for file %s: %fdB\n", args[0], rms)
	},
}

func init() {
	rootCmd.AddCommand(rmsCmd)
}

func GetOverallRMS(filePath string) (float64, error) {
	var audio *mp3.Mp3

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetOverallRMS(), nil
}