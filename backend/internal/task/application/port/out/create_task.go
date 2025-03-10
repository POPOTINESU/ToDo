//go:generate mockgen -source=create_task.go -typed=true -destination=./mocks/create_task_mock.go

package out

import (
	"ToDo/pkg"
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

type CreateTaskPort[T pkg.Tx] interface {
	Create(ctx context.Context, dto *CreateTaskDTO, tx T) (uuid.UUID, error)
}
