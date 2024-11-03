package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	"github.com/natserract/toktik/pkg/text_processor"
	"github.com/natserract/toktik/shared/util"
)

type CreateUserInterestEmbeddingHandler struct {
	inMemoryRepository repositories.UserInterestsEmbeddingRepository
}

func NewCreateUserInterestEmbeddingHandler(r repositories.UserInterestsEmbeddingRepository) *CreateUserInterestEmbeddingHandler {
	return &CreateUserInterestEmbeddingHandler{
		inMemoryRepository: r,
	}
}

const (
	// If dev true, embeddings will load from local (sample/)
	// Otherwise, generated from LLM
	__DEV__ = true
)

func (c *CreateUserInterestEmbeddingHandler) Handle(
	ctx context.Context,
	query CreateUserInterestEmbedding,
) error {
	// Only in dev mode
	// Query: `Indonesia`
	if __DEV__ {
		// Read json data
		file, err := os.Open("sample/data.json")
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()

		byteValue, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		// Unmarshal the JSON data from the file
		var data []repositories.SaveUserInterestsEmbeddingModel
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON from file: %v", err)
		}

		// Store to cache
		if err := c.inMemoryRepository.SaveUserInterestsEmbedding(query.Actor, &data); err != nil {
			return err
		}
		return nil
	}

	// If already cached, doesn't need to create embeddings
	_, err := c.inMemoryRepository.GetUserInterestsEmbedding(query.Actor)
	if err != nil {
		var models []repositories.SaveUserInterestsEmbeddingModel
		embed := embedding.NewVectorEmbedding()

		// Production mode
		// Clean text
		cp := text_processor.CleanProcessor{}
		processRule := map[string]interface{}{
			"rules": map[string]interface{}{
				"pre_processing_rules": []interface{}{
					map[string]interface{}{
						"id":      "remove_extra_spaces",
						"enabled": true,
					},
					map[string]interface{}{
						"id":      "remove_urls_emails",
						"enabled": true,
					},
				},
			},
		}

		for _, content := range query.PageContents {
			text := cp.Clean(content, processRule)
			if text == "" {
				fmt.Println("error in clean text")
			}

			// Split into 'Tags', 'Title', & 'Author'
			textSplitted := c.textSplitter(text)
			textSplitted.Title = util.MaxSubstring(textSplitted.Title, 500)

			// Embeddings
			var tagsVec []float32
			var titleVec []float32
			var err error

			if textSplitted.Tags != "" {
				tagsVec, err = embed.CreateVector(textSplitted.Tags, ctx)
				if err != nil {
					log.Fatalf("Error creating tags embedding: %v", err)
				}
			}

			if textSplitted.Title != "" {
				titleVec, err = embed.CreateVector(textSplitted.Title, ctx)
				if err != nil {
					log.Fatalf("Error creating title embedding: %v", err)
				}
			}

			models = append(models, repositories.SaveUserInterestsEmbeddingModel{
				ID:          query.Actor,
				Tags:        textSplitted.Tags,
				TagsVector:  tagsVec,
				Title:       textSplitted.Title,
				TitleVector: titleVec,
			})
		}

		// Store to the cache
		if err := c.inMemoryRepository.SaveUserInterestsEmbedding(query.Actor, &models); err != nil {
			return err
		}
	}

	return nil
}

func (c *CreateUserInterestEmbeddingHandler) FindMostSimilar(queryVector []float32, models []repositories.SaveUserInterestsEmbeddingModel, vectorType string) (string, float32, error) {
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

func (c *CreateUserInterestEmbeddingHandler) textSplitter(input string) CreateUserInterestEmbeddingMetadata {
	var authors []string
	var tags []string
	var title []string

	words := strings.Fields(input)

	for _, word := range words {
		word = strings.ToLower(word)
		if strings.HasPrefix(word, "@") {
			authors = append(authors, word)
		} else if strings.HasPrefix(word, "#") {
			tags = append(tags, word)
		} else {
			title = append(title, word)
		}
	}

	return CreateUserInterestEmbeddingMetadata{
		Tags:  strings.Join(tags, " "),
		Title: strings.Join(title, " "),
	}
}
