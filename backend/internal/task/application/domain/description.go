package domain

import (
	"errors"
	"fmt"
)

const MaxDescriptionLength = 255

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

	if len(description) > MaxDescriptionLength {
		return fmt.Errorf(fmt.Sprintf("description must be at most %d characters", MaxDescriptionLength))
	}

	return nil
}

func (v Description) Value() string {
	return v.value
}
