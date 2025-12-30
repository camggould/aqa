package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Mp3Operation func(*cobra.Command, []string, string) string

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
	rootCmd.PersistentFlags().String("mp3", "", "the mp3 file to analyze.")
}

func HandleAudioAnalysis(cmd *cobra.Command, args []string, operation Mp3Operation) string {
	mp3File, err := cmd.Flags().GetString("mp3")

	if err != nil || mp3File == "" {
		return fmt.Sprintf("Failed to read mp3 flag: %s\n", err)
	}

	return operation(cmd, args, mp3File)
}