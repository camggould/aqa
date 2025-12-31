package utils

import "fmt"

func PrintDb(input float64) string {
	return fmt.Sprintf("%.2f dB", input)
}

func PrintSampleRate(input int) string {
	return fmt.Sprintf("%.1fk", float64(input) / 1000)
}