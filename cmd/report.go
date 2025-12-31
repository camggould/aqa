package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/templates"
	"github.com/camggould/aqa/utils"

	"path/filepath"

	"github.com/spf13/cobra"
)

type ReportResponse struct {
	File string `json:"file"`
	OutputDirectory string `json:"outputDir"`
}

var reportCmd = &cobra.Command{
	Use: "report",
	Short: "Generate an HTML report of audio quality for the provided audio file.",
	Long: "Generate an HTML report of audio quality for the provided audio file. The first argument is the input audio file, and the second argument is an optional output file location for the report.",
	Run: RunGenerator(runReportCmd),
}

func runReportCmd(cmd *cobra.Command, args []string, audioFile string) string {
	outfile, err := cmd.Flags().GetString("o")
	reportPath, err := GenerateReport(audioFile, outfile)

	if err != nil {
		return fmt.Sprintf("Report could not be generated due to error: %s\n", err)
	}

	responseData := ReportResponse{
		File: audioFile,
		OutputDirectory: reportPath,
	}

	return utils.FormattedJsonOutput(responseData)
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
	channels := audio.GetChannelCount()

	var outfile string

	if outputDir != "" {
		outfile = outputDir
	} else {
		outfile = "report.html"
	}

	fullPath := filepath.Dir(outfile)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return "", err
	}

	f, err := os.Create(outfile)

	if err != nil { return "", err }

	err = templates.Report(filePath, overallRms, rmsFloor, peakLevel, channels).Render(context.Background(), f)

	if err != nil {
		return "", fmt.Errorf("failed to write output file: %w", err)
	}

	return outfile, nil
}