# AQA: Audio Quality Assurance

AQA (pronounced "aqua") is a command line utility for evaluating the quality of audio files, particularly for professional narration or voiceovers.

## Quickstart

Clone the repo to start using aqa. Sample commands:

- `aqa peak --mp3 /path/to/my/mp3` returns the peak level for the MP3 file.
- `aqa rms --mp3 /path/to/my/mp3` returns the RMS for the MP3 file.
- `aqa rmsFloor --mp3 /path/to/my/mp3` returns the RMS floor of the audio file.
- `aqa rmsCeiling --mp3 /path/to/my/mp3` returns the RMS ceiling of the audio file.
- `aqa report --mp3 /path/to/my/mp3 --o my_output_file.html` generates an HTML report highlighting audio quality.

## Example report

<img width="839" height="203" alt="image" src="https://github.com/user-attachments/assets/ead3b603-1ef7-4879-ae9e-3c99ed2bb533" />

## Supported audio formats

Currently the only supported audio format is MP3.

## Contributing

Feel free to put out a PR for review if you'd like to add additional tools.

## About the tool

I am a former developer for ACX.com. I enjoyed the space and learned a lot about audio along the way. This is a hobby project to bring more tooling to the space over time. The tool itself utilizes ffmpeg-go and go-ffprobe under the hood for interfacing with audio files.
