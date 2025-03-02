package in

import (
	"context"

	"github.com/google/uuid"
)

type CreateTaskDTO struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Priority string `json:"priority"`
	Stage string `json:"stage"`
}

type ICreateTaskUseCase interface {
	Execute(ctx context.Context, dto *CreateTaskDTO) (uuid.UUID, error)
}