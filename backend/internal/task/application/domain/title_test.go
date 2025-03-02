package domain_test

import (
	"fmt"
	"strings"
	"testing"

	"ToDo/internal/task/application/domain"
)

func TestNewTitle(t *testing.T) {
	var maxTitle = strings.Repeat("a", domain.MaxTitleLength)

	data := []struct {
		testName string
		value    string
		expected string
		errMsg   string
	}{
		{"create title object", "test title", "test title", ""},
		{"cannot create empty title", "", "", "title must not be empty"},
		{"too long title", maxTitle + "a", "", fmt.Sprintf("title must be at most %d characters", domain.MaxTitleLength)},
		{"max length title", maxTitle, maxTitle, ""},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewTitle(d.value)

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
