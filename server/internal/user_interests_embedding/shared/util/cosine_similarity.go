package util

import (
	"fmt"
	"sort"

	"github.com/khaibin/go-cosinesimilarity"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	"github.com/natserract/toktik/shared/util"
)

func FindMostSimilar(queryVector []float32, models []repositories.SaveUserInterestsEmbeddingModel, threshold float64, vectorType string) ([]string, []float64, error) {
	var similarityContent []string
	var similarityScores []float64

	queryVector64 := util.Float32ToFloat64(queryVector)
	for _, data := range models {
		switch vectorType {
		case "tags":
			tagVector := util.Float32ToFloat64(data.TagsVector)

			// Ensure both vectors are non-empty and of the same length
			if len(queryVector) == 0 || len(tagVector) == 0 || len(queryVector) != len(tagVector) {
				continue
			}

			// Compute cosine similarity
			similarityMatrix := cosinesimilarity.Compute(
				[][]float64{queryVector64},
				[][]float64{tagVector},
			)
			if len(similarityMatrix) > 0 && len(similarityMatrix[0]) > 0 {
				similarity := similarityMatrix[0][0]
				if similarity >= threshold {
					similarityContent = append(similarityContent, data.Tag)
					similarityScores = append(similarityScores, similarity)
				}
			}
		case "title":
			titleVector := util.Float32ToFloat64(data.TitleVector)

			// Ensure both vectors are non-empty and of the same length
			if len(queryVector) == 0 || len(titleVector) == 0 || len(queryVector) != len(titleVector) {
				continue
			}

			// Compute cosine similarity
			similarityMatrix := cosinesimilarity.Compute(
				[][]float64{queryVector64},
				[][]float64{titleVector},
			)
			if len(similarityMatrix) > 0 && len(similarityMatrix[0]) > 0 {
				similarity := similarityMatrix[0][0]
				if similarity >= threshold {
					similarityContent = append(similarityContent, data.Title)
					similarityScores = append(similarityScores, similarity)
				}
			}
		default:
			return nil, nil, fmt.Errorf("unknown vector type: %s", vectorType)
		}
	}

	similarityContent, similarityScores = sortSimilaritiesDescending(similarityContent, similarityScores)
	return similarityContent, similarityScores, nil
}

type similarityPair struct {
	content string
	score   float64
}

func sortSimilaritiesDescending(similarityContent []string, similarityScores []float64) ([]string, []float64) {
	// Create a slice of pairs
	pairs := make([]similarityPair, len(similarityScores))
	for i := range similarityScores {
		pairs[i] = similarityPair{content: similarityContent[i], score: similarityScores[i]}
	}

	// Sort the pairs by score in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	// Extract the sorted content and scores
	for i, pair := range pairs {
		similarityContent[i] = pair.content
		similarityScores[i] = pair.score
	}

	return similarityContent, similarityScores
}
