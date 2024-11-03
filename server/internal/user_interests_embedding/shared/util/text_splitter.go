package util

import (
	"strings"

	"github.com/natserract/toktik/internal/user_interests_embedding"
)

func TextSplitter(input string) user_interests_embedding.CreateUserInterestEmbeddingMetadata {
	var authors []string
	var tags []string
	var titles []string

	words := strings.Fields(input)

	for _, word := range words {
		if strings.HasPrefix(word, "@") {
			authors = append(authors, word)
		} else if strings.HasPrefix(word, "#") {
			tags = append(tags, word)
		} else {
			titles = append(titles, word)
		}
	}

	return user_interests_embedding.CreateUserInterestEmbeddingMetadata{
		Tags:   tags,
		Titles: titles,
	}
}
