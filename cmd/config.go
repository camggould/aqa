package cmd

func init() {
	configureFlags()
	configureCommands()
}

func configureFlags() {
	rootCmd.PersistentFlags().String("file", "", "the audio file to analyze.")
	reportCmd.Flags().String("o", "", "the html file to output a report to. Default is 'report.html'.")
}

/* Configures enabled commands.
 * Any time a new command is created, it must be added here to be accessible via CLI.
*/
func configureCommands() {
	rootCmd.AddCommand(channelsCmd)
	rootCmd.AddCommand(rmsCmd)
	rootCmd.AddCommand(peakLevelCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(rmsCeilingCmd)
	rootCmd.AddCommand(rmsFloorCmd)
	rootCmd.AddCommand(sampleRateCmd)
}