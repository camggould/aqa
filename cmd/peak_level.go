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
		fmt.Printf(HandleMp3Analysis(cmd, args, runPeakLevelCommand))
		return
	},
}

func runPeakLevelCommand(cmd *cobra.Command, args []string, mp3File string) string {
	rms, err := GetPeakLevel(mp3File)

	if err != nil {
		return fmt.Sprintf("Peak level could not be calculated due to error: %s\n", err)
	}

	return fmt.Sprintf("Peak level for file %s: %fdB\n", mp3File, rms)
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