package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SearchFeeds struct {
	Keywords string
	Count    string
}

func (s *SearchFeeds) Validate() error {
	return validation.ValidateStruct(s, validation.Field(&s.Keywords, validation.Required))
}
