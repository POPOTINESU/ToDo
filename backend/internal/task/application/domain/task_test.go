package domain_test

import (
	"ToDo/internal/task/application/domain"
	"ToDo/pkg"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	taskID, err := uuid.NewV7()
	assert.NoError(t, err, "Failed to generate UUIDv7")

	data := []struct {
		name        string
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
		t.Run(d.name, func(t *testing.T) {
			task, err := domain.NewTask(d.id, d.title, d.description, d.priority, d.stage)

			if d.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, d.title, task.Title().Value())
				assert.Equal(t, d.description, task.Description().Value())
				assert.Equal(t, d.priority, task.Priority().Value())
				assert.Equal(t, d.stage, task.Stage().Value())
			}
		})
	}
}

func TestEqual(t *testing.T) {
	task1, _ := domain.NewTask(
		nil,
		"title",
		"description",
		domain.PriorityLow,
		domain.StageTodo,
	)

	task2, _ := domain.NewTask(
		nil,
		"title",
		"description",
		domain.PriorityLow,
		domain.StageTodo,
	)

	t.Run("Same task", func(t *testing.T) {
		assert.True(t, task1.Equal(&task1))
	})

	t.Run("Not same task", func(t *testing.T) {
		assert.False(t, task1.Equal(&task2))
	})
}
