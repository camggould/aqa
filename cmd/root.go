package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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