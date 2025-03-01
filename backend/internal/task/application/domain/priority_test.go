package domain_test

import (
	"ToDo/internal/task/application/domain"
	"testing"
)

func TestNewPriority(t *testing.T) {
	data := []struct {
		testName string
		value string
		expected string
		errMsg string
	}  {
		{"create low priority object", "low", "low", ""},
		{"create middle priority object", "low", "low", ""},
		{"create high object", "high", "high", ""},
		{"cannot create empty priority", "", "", "priority must not be empty"},
		{"unexpected word", "unexpected priority", "", "unexpected priority value"},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewPriority(d.value)

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