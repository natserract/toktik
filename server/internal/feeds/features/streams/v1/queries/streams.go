package queries

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Streams struct {
	Id string
}

func (s *Streams) Validate() error {
	return validation.ValidateStruct(s, validation.Field(&s.Id, validation.Required))
}
