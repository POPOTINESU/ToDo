package domain_test

import (
	"ToDo/internal/task/application/domain"
	"testing"
)

func TestNewStage(t * testing.T) {
	data := []struct {
		testName string
		value string
		expected string
		errMsg string
	} {
		{"create TODO stage", domain.StageTodo, domain.StageTodo, ""},
		{"create In Progress stage", domain.StageInProgress, domain.StageInProgress, ""},
		{"create Done", domain.StageDone, domain.StageDone, ""},
		{"cannot create empty stage", "", "", "stage must not be empty"},
		{"unexpected word", "unexpected stage", "", "unexpected stage value"},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			result, err := domain.NewStage(d.value)

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