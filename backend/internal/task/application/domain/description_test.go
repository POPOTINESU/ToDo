package domain_test

import (
	"ToDo/internal/task/application/domain"
	"fmt"
	"strings"
	"testing"
)

func TestNewDescription(t *testing.T) {
	var maxDescription = strings.Repeat("a", domain.MAX_DESCRIPTION_LENGTH)

	data := []struct {
		testName string
		value    string
		expected string
		errMsg   string
	}{
		{"Create description object", "test description", "test description", ""},
		{"Cannot create empty description", "", "", "description must not be empty"},
		{"Too long description", maxDescription + "a", "", fmt.Sprintf("description must be at most %d characters", domain.MAX_DESCRIPTION_LENGTH)},
		{"Max length description", maxDescription, maxDescription, ""},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewDescription(d.value)

			if d.expected != result.Value() {
				t.Errorf("期待する結果 %s, 得られた結果 %s", d.expected, result.Value())
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if d.errMsg != errMsg {
				t.Errorf("予想されたエラーメッセージ %s, 得られたメッセージ %s", d.errMsg, errMsg)
			}
		})
	}
}
