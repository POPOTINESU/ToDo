//go:generate mockgen -source=create_task.go -typed=true -destination=./mocks/create_task_mock.go
package in

import (
	"context"

	"github.com/google/uuid"
)

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Stage       string `json:"stage"`
}

type CreateTaskUseCase interface {
	Execute(ctx context.Context, dto *CreateTaskDTO) (uuid.UUID, error)
}
