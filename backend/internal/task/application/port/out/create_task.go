package out

import (
	"context"

	"github.com/google/uuid"
)

type CreateTaskDTO struct {
	ID          uuid.UUID
	Title       string
	Description string
	Priority    string
	Stage       string
}

type ICreateTaskPort[T any] interface {
	Create(ctx context.Context, dto *CreateTaskDTO, tx T) (uuid.UUID, error)
}