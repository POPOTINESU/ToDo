package domain

import (
	"errors"
	"fmt"
)

const MAX_TITLE_LENGTH = 20

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

	if len(title) > MAX_TITLE_LENGTH {
		return fmt.Errorf(fmt.Sprintf("title must be at most %d characters", MAX_TITLE_LENGTH))
	}

	return nil
}

func (t Title) Value() string {
	return t.value
}