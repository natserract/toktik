package unstructured

import (
	"regexp"
)

func RemoveWhiteSpace(str string) string {
	re := regexp.MustCompile(`\s+`)
	cleanedStr := re.ReplaceAllString(str, "")
	return cleanedStr
}
