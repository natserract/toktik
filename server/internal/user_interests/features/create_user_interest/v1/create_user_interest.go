package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUserMetadata struct {
	VideoID    string
	Title      string
	AuthorName string
	CreateTime int64
}

type CreateUserInterest struct {
	PageContent string
	Metadata    CreateUserMetadata
}

func (s *CreateUserInterest) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.PageContent, validation.Required),
		validation.Field(&s.Metadata, validation.Required),
	)
}
