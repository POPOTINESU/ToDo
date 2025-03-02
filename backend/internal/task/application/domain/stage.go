package domain

import "errors"

const (
	StageTodo       = "TODO"
	StageInProgress = "IN_PROGRESS"
	StageDone       = "Done"
)

type Stage struct {
	value string
}

func NewStage(stage string) (Stage, error) {
	if err := validateStage(stage); err != nil {
		return Stage{}, err
	}

	return Stage{value: stage}, nil
}

func validateStage(stage string) error {
	switch stage {
	case StageTodo, StageInProgress, StageDone:
		return nil
	case "":
		return errors.New("stage must not be empty")
	default:
		return errors.New("unexpected stage value")
	}
}

func (s Stage) Value() string {
	return s.value
}
