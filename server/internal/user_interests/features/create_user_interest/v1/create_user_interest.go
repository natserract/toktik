package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUserMetadata struct {
	Tags  string
	Title string
}

type CreateUserInterest struct {
	Actor        string
	PageContents []string
}

func (s *CreateUserInterest) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Actor, validation.Required),
		validation.Field(&s.PageContents, validation.Required),
	)
}
