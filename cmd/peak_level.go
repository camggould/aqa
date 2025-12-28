package cmd

import (
	"fmt"

	"github.com/camggould/aqa/mp3"

	"github.com/spf13/cobra"
)

var peakLevelCmd = &cobra.Command{
	Use: "peak",
	Short: "Get the peak level of an mp3 file.",
	Long: "Get the peak level of a provided mp3 file. Must be a valid MP3.",
	Run: func(cmd *cobra.Command, args []string) {
		mp3File, err := cmd.Flags().GetString("mp3")

		if err != nil {
			fmt.Printf("Failed to read mp3 flag: %s\n", err)
			return
		}
		rms, err := GetPeakLevel(mp3File)

		if err != nil {
			fmt.Printf("Peak level could not be calculated due to error: %s\n", err)
			return
		}

		fmt.Printf("Peak level for file %s: %fdB\n", mp3File, rms)
	},
}

func init() {
	rootCmd.AddCommand(peakLevelCmd)
}

func GetPeakLevel(filePath string) (float64, error) {
	var audio *mp3.Mp3

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetPeakDBFS(), nil
}