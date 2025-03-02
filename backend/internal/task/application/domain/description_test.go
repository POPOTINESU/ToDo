package domain_test

import (
	"ToDo/internal/task/application/domain"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDescription(t *testing.T) {
	var maxDescription = strings.Repeat("a", domain.MaxDescriptionLength)

	data := []struct {
		testName string
		value    string
		expected string
		errMsg   string
	}{
		{"Create description object", "test description", "test description", ""},
		{"Cannot create empty description", "", "", "description must not be empty"},
		{"Too long description", maxDescription + "a", "", fmt.Sprintf("description must be at most %d characters", domain.MaxDescriptionLength)},
		{"Max length description", maxDescription, maxDescription, ""},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewDescription(d.value)

			assert.Equal(t, d.expected, result.Value(), "Unexpected description value")

			if d.errMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, d.errMsg)
			}
		})
	}
}
