package v1

import (
	"context"
	"fmt"

	"github.com/natserract/toktik/internal/user_interests/data/repositories"
)

type CreateUserInterestHandler struct {
	inMemoryRepository repositories.UserInterestsRepository
}

func NewCreateUserInterestHandler(r repositories.UserInterestsRepository) *CreateUserInterestHandler {
	return &CreateUserInterestHandler{
		inMemoryRepository: r,
	}
}

const (
	// # Temporary solution
	//
	// User interest cache stored by User Id / Name
	// At this time, I don't provided user authentication mechanism for  a while
	// So, just put my name as a User
	Natserract = "natserract"
)

func (c *CreateUserInterestHandler) Handle(
	ctx context.Context,
	query *CreateUserInterest,
) error {
	if err := c.inMemoryRepository.SaveUserInterest(Natserract, query.PageContent, query.Metadata); err != nil {
		return err
	}

	iterator := c.inMemoryRepository.DB().Cache.Iterator()
	for iterator.SetNext() {
		current, err := iterator.Value()
		if err != nil {
			return err
		}

		userInterests, err := c.inMemoryRepository.GetUserInterests(current.Key())
		if err != nil {
			return err
		}

		// Preprocessing text
		// Embedding
		// Find feed in cache
		fmt.Println("userInterests", userInterests)
	}

	return nil
}
