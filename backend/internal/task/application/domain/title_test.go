package domain_test

import (
	"ToDo/internal/task/application/domain"
	"testing"
)

func TestNewTitle(t *testing.T) {
	data := []struct {
		testName string
		value    string
		expected string
		errMsg   string
	}{
		{"create title object", "test title", "test title", ""},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err :=  domain.NewTitle(d.value)

			if result.Value() != d.expected {
				t.Errorf("期待する結果 %s, 得られた結果 %s", d.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if errMsg != d.errMsg {
				t.Errorf("予想されたエラーメッセージ %s, 得られたメッセージ %s", d.errMsg, errMsg)
			}
		})
	}
}