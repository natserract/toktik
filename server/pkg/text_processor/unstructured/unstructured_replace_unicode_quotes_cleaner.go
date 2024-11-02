package unstructured

import (
	"strings"
)

func ReplaceUnicodeQuotes(s string) string {
	replacer := strings.NewReplacer(
		"“", "\"", // Left double quotation mark
		"”", "\"", // Right double quotation mark
		"‘", "'", // Left single quotation mark
		"’", "'", // Right single quotation mark
		"«", "\"", // Left-pointing double angle quotation mark
		"»", "\"", // Right-pointing double angle quotation mark
		"‹", "'", // Single left-pointing angle quotation mark
		"›", "'", // Single right-pointing angle quotation mark
		"„", "\"", // Double low-9 quotation mark
		"‟", "\"", // Double high-reversed-9 quotation mark
		"‚", "'", // Single low-9 quotation mark
		"‛", "'", // Single high-reversed-9 quotation mark
	)

	return replacer.Replace(s)
}
