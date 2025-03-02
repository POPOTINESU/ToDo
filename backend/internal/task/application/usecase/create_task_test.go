package usecase_test

import (
	"ToDo/internal/task/application/domain"
	"ToDo/internal/task/application/port/in"
	mock_out "ToDo/internal/task/application/port/out/mocks"
	"ToDo/internal/task/application/usecase"
	mock_pkg "ToDo/pkg/mocks"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTaskExecute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTxManager := mock_pkg.NewMockTxManager[*mock_pkg.MockTx](ctrl)
	mockRepo := mock_out.NewMockCreateTaskPort[*mock_pkg.MockTx](ctrl)
	service := usecase.NewCreateTaskService(mockRepo, mockTxManager)

	ctx := context.Background()
	mockTx := mock_pkg.NewMockTx(ctrl)

	validDTO := &in.CreateTaskDTO{
		Title:       "Valid Title",
		Description: "Valid Description",
		Priority:    domain.PriorityHigh,
		Stage:       domain.StageTodo,
	}

	taskID := uuid.New()

	tests := []struct {
		name        string
		setupMocks  func()
		expectedErr error
	}{
		{
			name: "Successfully create task",
			setupMocks: func() {
				mockTxManager.EXPECT().Begin(ctx).Return(mockTx, nil)
				mockRepo.EXPECT().Create(ctx, gomock.Any(), mockTx).Return(taskID, nil)
				mockTx.EXPECT().Commit().Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Fail to begin transaction",
			setupMocks: func() {
				mockTxManager.EXPECT().Begin(ctx).Return(nil, errors.New("failed to begin transaction"))
			},
			expectedErr: errors.New("failed to begin transaction"),
		},
		{
			name: "Fail to create task (DB error) and rollback",
			setupMocks: func() {
				mockTxManager.EXPECT().Begin(ctx).Return(mockTx, nil)
				mockRepo.EXPECT().Create(ctx, gomock.Any(), mockTx).Return(uuid.Nil, errors.New("db error"))
				mockTx.EXPECT().Rollback().Return(nil)
			},
			expectedErr: errors.New("db error"),
		},
		{
			name: "Fail to commit transaction",
			setupMocks: func() {
				mockTxManager.EXPECT().Begin(ctx).Return(mockTx, nil)
				mockRepo.EXPECT().Create(ctx, gomock.Any(), mockTx).Return(taskID, nil)
				mockTx.EXPECT().Commit().Return(errors.New("commit error"))
				mockTx.EXPECT().Rollback().Return(nil)
			},
			expectedErr: errors.New("commit error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			id, err := service.Execute(ctx, validDTO)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Equal(t, uuid.Nil, id)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, taskID, id)
			}
		})
	}
}
