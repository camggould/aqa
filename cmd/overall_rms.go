package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/utils"
	"github.com/spf13/cobra"
)

type RmsResponse struct {
	File string `json:"file"`
	Rms string `json:"RMS"`
}

var rmsCmd = &cobra.Command{
	Use: "rms",
	Short: "Get the RMS of an audio file.",
	Long: "Get the RMS of a provided audio file. Must be a valid audio file.",
	Run: RunGenerator(runRmsCommand),
}

func runRmsCommand(cmd *cobra.Command, args []string, audioFile string) string {
	rms, err := GetOverallRMS(audioFile)
	if err != nil { return fmt.Sprintf("RMS could not be calculated due to error: %s\n", err) }

	responseData := RmsResponse{
		File: audioFile,
		Rms: utils.PrintDb(rms),
	}

	return utils.FormattedJsonOutput(responseData)
}

func GetOverallRMS(filePath string) (float64, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetOverallRMS(), nil
}