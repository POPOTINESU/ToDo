package domain_test

import (
	"ToDo/internal/task/application/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPriority(t *testing.T) {
	data := []struct {
		testName string
		value    string
		expected string
		errMsg   string
	}{
		{"Create low priority object", domain.PriorityLow, domain.PriorityLow, ""},
		{"Create middle priority object", domain.PriorityMiddle, domain.PriorityMiddle, ""},
		{"Create high object", domain.PriorityHigh, domain.PriorityHigh, ""},
		{"Cannot create empty priority", "", "", "priority must not be empty"},
		{"Unexpected word", "unexpected priority", "", "unexpected priority value"},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewPriority(d.value)

			assert.Equal(t, d.expected, result.Value(), "Unexpected priority value")

			if d.errMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, d.errMsg)
			}
		})
	}
}
