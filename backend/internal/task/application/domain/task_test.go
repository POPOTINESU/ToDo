package domain_test

import (
	"ToDo/internal/task/application/domain"
	"ToDo/pkg"
	"testing"

	"github.com/google/uuid"
)

func TestNewTask(t *testing.T) {
	taskID, err := uuid.NewV7()
	if err != nil {
		t.Fatalf("Failed to generate UUIDv7: %v", err) 
	}
	data := []struct {
		testName    string
		id          *uuid.UUID
		title       string
		description string
		priority    string
		stage       string
		expectErr   bool
	}{
		{"Create new task", nil, "Valid Title", "Valid Desc", domain.PriorityHigh, domain.StageTodo, false},
		{"Create exist task object", pkg.UUIDPtr(taskID), "Valid Title", "Valid Desc", domain.PriorityHigh, domain.StageTodo, false},
		{"Empty title", nil, "", "Valid Desc", domain.PriorityHigh, domain.StageTodo, true},
		{"Empty description", nil, "Valid Title", "", domain.PriorityHigh, domain.StageTodo, true},
		{"Empty priority", nil, "Valid Title", "Valid Desc", "", domain.StageTodo, true},
		{"Empty stage", nil, "Valid Title", "Valid Desc", domain.PriorityHigh, "", true},
	}

	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			task, err := domain.NewTask(d.id, d.title, d.description, d.priority, d.stage)

			if d.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else {
					t.Logf("unexpected error %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if task.Title().Value() != d.title {
					t.Errorf("Expected title %s, got %s", d.title, task.Title().Value())
				}
				if task.Description().Value() != d.description {
					t.Errorf("Expected description %s, got %s", d.description, task.Description().Value())
				}
				if task.Priority().Value() != d.priority {
					t.Errorf("Expected priority %s, got %s", d.priority, task.Priority().Value())
				}
				if task.Stage().Value() != d.stage {
					t.Errorf("Expected stage %s, got %s", d.stage, task.Stage().Value())
				}
			}
		})
	}
}
