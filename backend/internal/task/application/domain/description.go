package domain

import (
	"errors"
	"fmt"
)

const MAX_DESCRIPTION_LENGTH = 255

type Description struct {
	value string
}

func NewDescription(description string) (Description, error) {
	if err := validateDescription(description); err != nil {
		return Description{}, err
	}

	return Description{
		value: description,
	}, nil
}

func validateDescription(description string) error {
	if description == "" {
		return errors.New("description must not be empty")
	}

	if len(description) > MAX_DESCRIPTION_LENGTH {
		return fmt.Errorf(fmt.Sprintf("description must be at most %d characters", MAX_DESCRIPTION_LENGTH))
	}

	return nil
}

func (v Description) Value() string {
	return v.value
}
