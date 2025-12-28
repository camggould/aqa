package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/camggould/aqa/mp3"
	"github.com/camggould/aqa/templates"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use: "report",
	Short: "Generate an HTML report of audio quality for the provided MP3.",
	Long: "Generate an HTML report of audio quality for the provided MP3.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reportPath, err := GenerateReport(args[0])

		if err != nil {
			fmt.Printf("Report could not be generated due to error: %s\n", err)
			return
		}

		fmt.Printf("Report for file %s: can be found in %s.\n", args[0], reportPath)
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}

func GenerateReport(filePath string) (string, error) {
	var audio *mp3.Mp3
	audio, err := audio.New(filePath)

	if err != nil {
		return "", fmt.Errorf("failed to handle audio file: %w", err)
	}

	overallRms := audio.GetOverallRMS()
	rmsFloor := audio.GetRmsFloor()

	f, err := os.Create("report.html")

	if err != nil {
		return "", fmt.Errorf("Failed to create report file")
	}

	err = templates.Report(filePath, overallRms, rmsFloor).Render(context.Background(), f)

	if err != nil {
		return "", fmt.Errorf("failed to write output file: %w", err)
	}

	return "report.html", nil
}