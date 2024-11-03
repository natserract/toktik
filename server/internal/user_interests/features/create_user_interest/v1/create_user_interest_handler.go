package v1

import (
	"context"

	"github.com/natserract/toktik/internal/user_interests/data/repositories"
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
	// If already cached, doesn't need to create embeddings
	_, err := c.inMemoryRepository.GetUserInterests(query.Actor)
	if err != nil {
		var pageContents []string

		for _, content := range query.PageContents {
			pageContents = append(pageContents, util.MaxSubstring(content, 500))
		}

		// Store to the cache
		if err := c.inMemoryRepository.SaveUserInterests(query.Actor, pageContents); err != nil {
			return err
		}
	}

	return nil
}
