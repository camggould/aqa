package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/utils"

	"github.com/spf13/cobra"
)

type RmsCeilingResponse struct {
	File string `json:"file"`
	RmsCeiling string `json:"rmsCeiling"`
}

var rmsCeilingCmd = &cobra.Command{
	Use: "rmsCeiling",
	Short: "Get the RMS ceiling of an audio file.",
	Long: "Get the RMS ceiling of a provided audio file. Must be a valid audio file. Calculated based on the loudest 0.5s segment of audio.",
	Run: RunGenerator(runRmsCeilingCommand),
}

func runRmsCeilingCommand(cmd *cobra.Command, args []string, audioFile string) string {
	rmsCeiling, err := GetRmsCeiling(audioFile)

	if err != nil {
		return fmt.Sprintf("RMS ceiling could not be calculated due to error: %s\n", err)
	}

	responseData := RmsCeilingResponse{
		File: audioFile,
		RmsCeiling: utils.PrintDb(rmsCeiling),
	}

	return utils.FormattedJsonOutput(responseData)
}

func GetRmsCeiling(filePath string) (float64, error) {
	var audio *audio.AudioFile

	audio, err := audio.New(filePath)

	if err != nil {
		return 0.0, fmt.Errorf("failed to handle audio file: %w", err)
	}

	return audio.GetRMSCeiling(), nil

}