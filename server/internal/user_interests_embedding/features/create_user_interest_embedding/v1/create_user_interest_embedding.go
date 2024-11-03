package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUserInterestEmbeddingMetadata struct {
	Tags  string
	Title string
}

type CreateUserInterestEmbedding struct {
	Actor        string
	PageContents []string
}

func (s *CreateUserInterestEmbedding) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Actor, validation.Required),
		validation.Field(&s.PageContents, validation.Required),
	)
}
