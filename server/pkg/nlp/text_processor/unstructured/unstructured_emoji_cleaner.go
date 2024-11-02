package unstructured

func isEmoji(r rune) bool {
	// Emojis are generally in the range of certain Unicode blocks.
	// This is a simple check and might not cover all emojis.
	return (r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
		(r >= 0x1F300 && r <= 0x1F5FF) || // Miscellaneous Symbols and Pictographs
		(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and Map Symbols
		(r >= 0x2600 && r <= 0x26FF) || // Miscellaneous Symbols
		(r >= 0x2700 && r <= 0x27BF) || // Dingbats
		(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental Symbols and Pictographs
		(r >= 0x1F1E6 && r <= 0x1F1FF) // Flags (iOS)
}

func RemoveEmojis(input string) string {
	output := make([]rune, 0, len(input))
	for _, r := range input {
		if !isEmoji(r) {
			output = append(output, r)
		}
	}
	return string(output)
}
