package queries

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetFeedById struct {
	Id string
}

func (s *GetFeedById) Validate() error {
	return validation.ValidateStruct(s, validation.Field(&s.Id, validation.Required))
}
