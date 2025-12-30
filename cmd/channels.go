package cmd

import (
	"fmt"

	"github.com/camggould/aqa/audio"
	"github.com/camggould/aqa/utils"
	"github.com/spf13/cobra"
)

type ChannelsResponse struct {
	File string `json:"file"`
	Channels int `json:"channels"`
}

var channelsCmd = &cobra.Command{
	Use: "channels",
	Short: "Get the number of channels for an audio file.",
	Long: "Get the number of channels for a provided audio file. Must be a valid audio file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(HandleAudioAnalysis(cmd, args, runChannelsCommand))
		return
	},
}

func runChannelsCommand(cmd *cobra.Command, args []string, audioFile string) string {
	channels, err := GetChannelCount(audioFile)

	if err != nil {
		return fmt.Sprintf("Number of channels could not be found: %s\n", err)
	}

	channelsResponse := ChannelsResponse{
		File: audioFile,
		Channels: channels,
	}

	return utils.FormattedJsonOutput(channelsResponse)
}

func init() {
	rootCmd.AddCommand(channelsCmd)
}

func GetChannelCount(filePath string) (int, error) {
	var audio *audio.AudioFile
	audio, err := audio.New(filePath)
	if err != nil { return 0, err }
	return audio.GetChannelCount(), nil
}