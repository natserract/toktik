package v1

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
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

func (c *CreateUserInterestHandler) Handle(
	ctx context.Context,
	query *CreateUserInterest,
) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	if err := c.inMemoryRepository.Save(uuid.String(), query.PageContent, query.Metadata); err != nil {
		return err
	}
	fmt.Println(" query.PageContent", query.PageContent, query.Metadata)

	return nil
}
