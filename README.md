# AQA: Audio Quality Assurance

AQA (pronounced "aqua") is a command line utility for evaluating the quality of audio files, particularly for professional narration or voiceovers.

## Quickstart

Clone the repo to start using aqa. Sample commands:

- `aqa peak --file /path/to/my/audio.mp3` returns the peak level for the audio file.
- `aqa rms --file /path/to/my/audio.wav` returns the RMS for the audio file.
- `aqa rmsFloor --file /path/to/my/audio.mp3` returns the RMS floor of the audio file.
- `aqa rmsCeiling --file /path/to/my/audio.flac` returns the RMS ceiling of the audio file.
- `aqa channels --file /path/to/my/audio.flac` returns the number of channels in the audio file.
- `aqa report --file /path/to/my/audio.aac --o my_output_file.html` generates an HTML report highlighting audio quality.
  - The `--o` flag also supports directories.

## Example report

<img width="839" height="203" alt="image" src="https://github.com/user-attachments/assets/ead3b603-1ef7-4879-ae9e-3c99ed2bb533" />

## Supported audio formats

- MP3
- WAV
- FLAC
- AAC

Other audio files supported by ffmpeg may work, but these four audio formats have been tested.

## Contributing

Feel free to put out a PR for review if you'd like to add additional tools.

## About the tool

I am a former developer for ACX.com. I enjoyed the space and learned a lot about audio along the way. This is a hobby project to bring more tooling to the space over time. The tool itself utilizes ffmpeg-go and go-ffprobe under the hood for interfacing with audio files.
