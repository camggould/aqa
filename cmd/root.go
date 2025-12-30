package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type AudioOperation func(*cobra.Command, []string, string) string

var rootCmd = &cobra.Command{
	Use: "aqa",
	Short: "aqa is a cli tool for checking audio quality for professional voiceover.",
	Long: "aqa is a cli tool for checking audio quality for professional voiceover. It checks RMS floor, RMS range, peak values, channel format, and CBR.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while executing aqa '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("file", "", "the audio file to analyze.")
}

func HandleAudioAnalysis(cmd *cobra.Command, args []string, operation AudioOperation) string {
	audioFile, err := cmd.Flags().GetString("file")

	if err != nil || audioFile == "" {
		return fmt.Sprintf("Failed to read file flag: %s\n", err)
	}

	return operation(cmd, args, audioFile)
}