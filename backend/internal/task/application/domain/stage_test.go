package domain_test

import (
	"ToDo/internal/task/application/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStage(t *testing.T) {
	data := []struct {
		testName string
		value    string
		expected string
		errMsg   string
	}{
		{"Create TODO stage", domain.StageTodo, domain.StageTodo, ""},
		{"Create In Progress stage", domain.StageInProgress, domain.StageInProgress, ""},
		{"Create Done", domain.StageDone, domain.StageDone, ""},
		{"Cannot create empty stage", "", "", "stage must not be empty"},
		{"Unexpected word", "unexpected stage", "", "unexpected stage value"},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewStage(d.value)

			assert.Equal(t, d.expected, result.Value(), "Unexpected stage value")

			if d.errMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, d.errMsg)
			}
		})
	}
}
