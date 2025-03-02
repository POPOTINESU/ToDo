package domain

import "errors"

type Priority struct {
	value string
}

const (
	PriorityLow    = "LOW"
	PriorityMiddle = "MIDDLE"
	PriorityHigh   = "HIGH"
)

func NewPriority(priority string) (Priority, error) {
	if err := validatePriority(priority); err != nil {
		return Priority{}, err
	}

	return Priority{priority}, nil
}

func validatePriority(priority string) error {
	switch priority {
	case PriorityLow, PriorityMiddle, PriorityHigh:
		return nil
	case "":
		return errors.New("priority must not be empty")
	default:
		return errors.New("unexpected priority value")
	}
}

func (p Priority) Value() string {
	return p.value
}
