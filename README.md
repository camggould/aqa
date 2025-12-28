# AQA: Audio Quality Assurance

AQA (pronounced "aqua") is a command line utility for evaluating the quality of audio files, particularly for professional narration or voiceovers.

## Quickstart

Clone the repo to start using aqa. Sample commands:

`aqa rmsFloor /path/to/my/mp3` returns the RMS floor of the audio file.

## Supported audio formats

Currently the only supported audio format is MP3.

## Contributing

Feel free to put out a PR for review if you'd like to add additional tools.

## About the tool

I am a former developer for ACX.com. I enjoyed the space and learned a lot about audio along the way. This is a hobby project to bring more tooling to the space over time. The tool itself utilizes ffmpeg-go and go-ffprobe under the hood for interfacing with audio files.