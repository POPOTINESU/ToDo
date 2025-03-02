package domain

import (
	"errors"
	"fmt"
)

const MaxTitleLength = 20

type Title struct {
	value string
}

func NewTitle(title string) (Title, error) {
	if err := validateTitle(title); err != nil {
		return Title{}, err
	}

	return Title{value: title}, nil
}

func validateTitle(title string) error {
	if title == "" {
		return errors.New("title must not be empty")
	}

	if len(title) > MaxTitleLength {
		return fmt.Errorf("title must be at most %d characters", MaxTitleLength)
	}

	return nil
}

func (t Title) Value() string {
	return t.value
}
