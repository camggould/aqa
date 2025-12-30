package utils

import (
	"encoding/json"
	"fmt"
)

func FormattedJsonOutput(data any) string {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil { return fmt.Sprintf("Error: failed to parse output.")}
	return fmt.Sprintf("%s\n", string(output))
}