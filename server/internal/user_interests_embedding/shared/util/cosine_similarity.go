package util

import (
	"fmt"
	"log"

	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
)

func FindMostSimilar(queryVector []float32, models []repositories.SaveUserInterestsEmbeddingModel, vectorType string) (string, float32, error) {
	var maxSimilarity float32
	var bestText string

	embed := embedding.NewVectorEmbedding()
	for _, model := range models {
		var similarity float32
		var err error

		switch vectorType {
		case "tags":
			similarity, err = embed.CosineSimilarity(queryVector, model.TagsVector)
			if err == nil && similarity > maxSimilarity {
				maxSimilarity = similarity
				bestText = model.Tags
			}
		case "title":
			similarity, err = embed.CosineSimilarity(queryVector, model.TitleVector)
			if err == nil && similarity > maxSimilarity {
				maxSimilarity = similarity
				bestText = model.Title
			}
		default:
			return "", 0, fmt.Errorf("unknown vector type: %s", vectorType)
		}

		if err != nil {
			log.Printf("Error calculating similarity: %v", err)
		}
	}

	return bestText, maxSimilarity, nil
}
