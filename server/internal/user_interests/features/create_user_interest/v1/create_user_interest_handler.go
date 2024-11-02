package v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/natserract/toktik/internal/user_interests/data/repositories"
	"github.com/natserract/toktik/pkg/nlp/text_processor"
	"github.com/natserract/toktik/shared/util"
)

type CreateUserInterestHandler struct {
	inMemoryRepository repositories.UserInterestsRepository
}

func NewCreateUserInterestHandler(r repositories.UserInterestsRepository) *CreateUserInterestHandler {
	return &CreateUserInterestHandler{
		inMemoryRepository: r,
	}
}

func (c *CreateUserInterestHandler) Handle(
	ctx context.Context,
	query CreateUserInterest,
) error {
	fmt.Println("Actor: ", query.Actor)

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

	text := cp.Clean(query.PageContent, processRule)
	if text == "" {
		fmt.Println("error in clean text")
	}

	// Split into 'Tags', 'Title', & 'Author'
	textSplitted := c.textSplitter(text)
	textSplitted.Title = util.MaxSubstring(textSplitted.Title, 500)
	if err := c.inMemoryRepository.SaveUserInterest(
		query.Actor,
		textSplitted.Title,
		CreateUserMetadata{
			Tags:   textSplitted.Tags,
			Title:  textSplitted.Title,
			Author: textSplitted.Author,
		},
	); err != nil {
		return err
	}

	fmt.Println("Len ", c.inMemoryRepository.DB().Cache.Len())

	return nil
}

func (c *CreateUserInterestHandler) textSplitter(input string) CreateUserMetadata {
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

	return CreateUserMetadata{
		Tags:   strings.Join(tags, " "),
		Title:  strings.Join(title, " "),
		Author: strings.Join(authors, " "),
	}
}
