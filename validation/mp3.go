package validation

import (
	"fmt"
	"io"
	"os"

	"github.com/tcolgate/mp3"
)

func IsMP3(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	d := mp3.NewDecoder(f)
	var frame mp3.Frame
	var skipped int

	if err := d.Decode(&frame, &skipped); err != nil {
		if err == io.EOF {
			return fmt.Errorf("file is empty or too short to be a valid MP3")
		}

		return fmt.Errorf("file is not a valid mp3: %w", err)
	}

	return nil
}