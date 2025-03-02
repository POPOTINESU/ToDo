package usecase

import (
	"ToDo/internal/task/application/domain"
	"ToDo/internal/task/application/port/in"
	"ToDo/internal/task/application/port/out"
	"ToDo/pkg"
	"context"

	"github.com/google/uuid"
)

// CreateTaskService handles task creation logic.
type CreateTaskService[T pkg.Tx] struct {
	repo      out.CreateTaskPort[T]
	txManager pkg.TxManager[T]
}

func NewCreateTaskService[T pkg.Tx](repo out.CreateTaskPort[T], txManager pkg.TxManager[T]) in.CreateTaskUseCase {
	return &CreateTaskService[T]{
		repo:      repo,
		txManager: txManager,
	}
}

// Execute creates a new task and handles the transaction.
func (s *CreateTaskService[T]) Execute(ctx context.Context, dto *in.CreateTaskDTO) (uuid.UUID, error) {
	tx, err := s.txManager.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	taskEntity, err := s.convertDTOToTaskEntity(dto)
	if err != nil {
		return uuid.Nil, err
	}

	outputDTO := s.convertEntityToOutputDTO(&taskEntity)
	createdTaskID, err := s.repo.Create(ctx, &outputDTO, tx)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	return createdTaskID, nil
}

func (s *CreateTaskService[T]) convertDTOToTaskEntity(dto *in.CreateTaskDTO) (domain.Task, error) {
	taskEntity, err := domain.NewTask(
		nil,
		dto.Title,
		dto.Description,
		dto.Priority,
		dto.Stage,
	)

	if err != nil {
		return domain.Task{}, err
	}

	return taskEntity, nil
}

func (s *CreateTaskService[T]) convertEntityToOutputDTO(entity *domain.Task) out.CreateTaskDTO {
	return out.CreateTaskDTO{
		ID:          entity.ID(),
		Title:       entity.Title().Value(),
		Description: entity.Description().Value(),
		Priority:    entity.Priority().Value(),
		Stage:       entity.Stage().Value(),
	}
}
