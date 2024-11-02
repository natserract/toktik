package unstructured

func RemoveNonASCII(s string) string {
	asciiOnly := make([]rune, 0, len(s))
	for _, r := range s {
		if r <= 127 { // ASCII characters are in the range 0 to 127
			asciiOnly = append(asciiOnly, r)
		}
	}
	return string(asciiOnly)
}
