package text_processor

import (
	"regexp"

	"github.com/natserract/toktik/pkg/nlp/text_processor/unstructured"
)

type CleanProcessor struct{}

func (cp *CleanProcessor) Clean(text string, processRule map[string]interface{}) string {
	// Default clean
	// Remove invalid symbols
	re := regexp.MustCompile(`<\|`)
	text = re.ReplaceAllString(text, "<")
	re = regexp.MustCompile(`\|>`)
	text = re.ReplaceAllString(text, ">")
	re = regexp.MustCompile(`[\x00-\x08\x0B\x0C\x0E-\x1F\x7F\xEF\xBF\xBE]`)
	text = re.ReplaceAllString(text, "")
	// Unicode U+FFFE
	re = regexp.MustCompile(`\x{FFFE}`)
	text = re.ReplaceAllString(text, "")

	// Remove emojis
	text = unstructured.RemoveEmojis(text)

	rules, ok := processRule["rules"].(map[string]interface{})
	if ok {
		if preProcessingRules, ok := rules["pre_processing_rules"].([]interface{}); ok {
			for _, rule := range preProcessingRules {
				if ruleMap, ok := rule.(map[string]interface{}); ok {
					if ruleMap["id"] == "remove_extra_spaces" && ruleMap["enabled"] == true {
						// Remove extra spaces
						re = regexp.MustCompile(`\n{3,}`)
						text = re.ReplaceAllString(text, "\n\n")
						re = regexp.MustCompile(`[\t\f\r\x20\x{00A0}\x{1680}\x{180E}\x{2000}-\x{200A}\x{202F}\x{205F}\x{3000}]{2,}`)
						text = re.ReplaceAllString(text, " ")
					} else if ruleMap["id"] == "remove_urls_emails" && ruleMap["enabled"] == true {
						// Remove email
						re = regexp.MustCompile(`([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`)
						text = re.ReplaceAllString(text, "")

						// Remove URL
						re = regexp.MustCompile(`https?://[^\s]+`)
						text = re.ReplaceAllString(text, "")
					}
				}
			}
		}
	}

	return text
}

func (cp *CleanProcessor) FilterString(text string) string {
	return text
}
