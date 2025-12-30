package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type AudioOperation func(*cobra.Command, []string, string) string
type AudioOperationHandler func(cmd *cobra.Command, args []string, operation AudioOperation)
type RunFunction func(*cobra.Command, []string)

var rootCmd = &cobra.Command{
	Use: "aqa",
	Short: "aqa is a cli tool for checking audio quality for professional voiceover.",
	Long: "aqa is a cli tool for checking audio quality for professional voiceover. It checks RMS floor, RMS range, peak values, channel format, and CBR.",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while executing aqa '%s'\n", err)
		os.Exit(1)
	}
}

/** Generic helper for handling audio analysis commands.
 * This helper takes in the audio operation to invoke and invokes it after doing basic validation on the required flag.
*/
func HandleAudioAnalysis(cmd *cobra.Command, args []string, operation AudioOperation) {
	audioFile, err := cmd.Flags().GetString("file")

	if err != nil || audioFile == "" {
		fmt.Sprintf("Failed to read file flag: %s\n", err)
	}

	fmt.Printf(operation(cmd, args, audioFile))
}

/** Generic helper for generating run functions.
 * This helper makes syntax for creating "Run" cobra functions simple by only requiring RunGenerator(func)
 * instead of repetitive use of `func (cmd *cobra.Command, args []string) { HandleAudioAnalysis(cmd, args, myOperation) }`
 * for each new command being created.
*/
func RunGenerator(operation AudioOperation) RunFunction {
	return func (cmd *cobra.Command, args []string) { HandleAudioAnalysis(cmd, args, operation) }
}