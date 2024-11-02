package util

import "unicode/utf8"

func MaxSubstring(input string, maxLength int) string {
	if maxLength <= 0 {
		return ""
	}

	// Check if the input string is shorter than the maxLength
	if utf8.RuneCountInString(input) <= maxLength {
		return input
	}

	// Create a slice of runes to handle Unicode characters correctly
	runes := []rune(input)

	// Return the substring with the specified maximum length
	return string(runes[:maxLength])
}
