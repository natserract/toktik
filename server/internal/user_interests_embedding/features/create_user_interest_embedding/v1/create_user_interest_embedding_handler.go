package v1

import (
	"context"
	"encoding/json"
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
	__DEV__ = false
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

	// Production mode
	// If already cached, doesn't need to create embeddings
	_, err := c.inMemoryRepository.GetUserInterestsEmbedding(query.Actor)
	if err != nil {
		var models []repositories.SaveUserInterestsEmbeddingModel
		embed := embedding.NewVectorEmbedding()

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

		// Embeddings
		for _, content := range query.PageContents {
			// Split into 'Tags', 'Title', & 'Author'
			textSplitted := c.textSplitter(content)

			// var err error

			// Tags
			cachedTagText := string("")
			for _, tag := range textSplitted.Tags {
				tagText := util.MaxSubstring(tag, 700)
				tagsNormalized := cp.Clean(tagText, processRule)
				if tagsNormalized != "" {
					tagText = strings.ToLower(tagsNormalized)
					cachedTagText += " " + tagText
				}
			}
			log.Println("Create tags embedding...", query.Actor, cachedTagText)
			var tagsVec []float32
			if cachedTagText != "" {
				tagsVec, err = embed.CreateVector(cachedTagText, ctx)
				if err != nil {
					log.Fatalf("Error creating tags embedding: %v", err)
				}
				log.Println("Embedding tags succesfully created")
			}

			// Titles
			cachedTitleText := string("")
			for _, title := range textSplitted.Titles {
				titleText := util.MaxSubstring(title, 700)
				titleNormalized := cp.Clean(titleText, processRule)
				if titleNormalized != "" {
					titleText = strings.ToLower(titleNormalized)
					cachedTitleText += " " + titleText
				}
			}
			var titleVec []float32
			log.Println("Create titles embedding...", query.Actor, cachedTitleText)
			if cachedTitleText != "" {
				titleVec, err = embed.CreateVector(cachedTitleText, ctx)
				if err != nil {
					log.Fatalf("Error creating title embedding: %v", err)
				}
				log.Println("Embedding title succesfully created")
			}

			models = append(models, repositories.SaveUserInterestsEmbeddingModel{
				ID:          query.Actor,
				Tag:         strings.Join(textSplitted.Tags, " "),
				TagsVector:  tagsVec,
				Title:       strings.Join(textSplitted.Titles, " "),
				TitleVector: titleVec,
			})
			log.Println("Embedding created succesfully")
		}

		// Store to the cache
		if err := c.inMemoryRepository.SaveUserInterestsEmbedding(query.Actor, &models); err != nil {
			return err
		}
	}

	return nil
}

func (c *CreateUserInterestEmbeddingHandler) textSplitter(input string) CreateUserInterestEmbeddingMetadata {
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

	return CreateUserInterestEmbeddingMetadata{
		Tags:   tags,
		Titles: titles,
	}
}
