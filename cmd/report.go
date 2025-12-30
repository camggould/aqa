package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/templates"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use: "report",
	Short: "Generate an HTML report of audio quality for the provided audio file.",
	Long: "Generate an HTML report of audio quality for the provided audio file. The first argument is the input audio file, and the second argument is an optional output file location for the report.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(HandleAudioAnalysis(cmd, args, runReportCmd))
		return
	},
}

func runReportCmd(cmd *cobra.Command, args []string, audioFile string) string {
	outfile, err := cmd.Flags().GetString("o")
	reportPath, err := GenerateReport(audioFile, outfile)

	if err != nil {
		return fmt.Sprintf("Report could not be generated due to error: %s\n", err)
	}

	return fmt.Sprintf("Report for file %s: can be found in %s.\n", audioFile, reportPath)
}

func init() {
	reportCmd.Flags().String("o", "example.html", "the html file to output a report to. Default is 'report.html'.")
	rootCmd.AddCommand(reportCmd)
}

func GenerateReport(filePath string, outputDir string) (string, error) {
	var audio *audio.AudioFile
	audio, err := audio.New(filePath)

	if err != nil {
		return "", fmt.Errorf("failed to handle audio file: %w", err)
	}

	overallRms := audio.GetOverallRMS()
	rmsFloor := audio.GetRmsFloor()
	peakLevel := audio.GetPeakDBFS()

	var outfile string

	if outputDir != "" {
		outfile = outputDir
	} else {
		outfile = "report.html"
	}

	f, err := os.Create(outfile)

	if err != nil {
		return "", fmt.Errorf("Failed to create report file")
	}

	err = templates.Report(filePath, overallRms, rmsFloor, peakLevel).Render(context.Background(), f)

	if err != nil {
		return "", fmt.Errorf("failed to write output file: %w", err)
	}

	return outfile, nil
}