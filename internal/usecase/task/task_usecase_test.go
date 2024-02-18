package task

import (
	"service-task-list/config"
	"service-task-list/internal/entity"
	"service-task-list/internal/repository/mocks"
	"service-task-list/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("CreateTask", &entity.TaskRequest{
		Title: "test",
	}).Return(nil).Once()

	// test
	err := taskUsecase.CreateTask(&entity.TaskRequest{
		Title: "test",
	})
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestCreateTaskError(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("CreateTask", &entity.TaskRequest{
		Title: "test",
	}).Return(assert.AnError).Once()

	// test
	err := taskUsecase.CreateTask(&entity.TaskRequest{
		Title: "test",
	})
	assert.Error(t, err)
	mockTaskRepo.AssertExpectations(t)

}

func TestGetTaskList(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("GetTaskList", 0).Return([]*entity.Task{
		{
			ID:          1,
			Title:       "test",
			Description: "test",
			Completed:   false,
		},
	}, nil).Once()

	// test
	tasks, err := taskUsecase.GetTaskList(0)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.Equal(t, 1, len(tasks))
	mockTaskRepo.AssertExpectations(t)
}

func TestGetTaskListError(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("GetTaskList", 0).Return(nil, assert.AnError).Once()

	// test
	tasks, err := taskUsecase.GetTaskList(0)
	assert.Error(t, err)
	assert.Nil(t, tasks)
	mockTaskRepo.AssertExpectations(t)
}

func TestCheckTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("CheckTask", 1).Return(nil).Once()

	// test
	err := taskUsecase.CheckTask(1)
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("UpdateTask", &entity.Task{
		ID:          1,
		Title:       "test",
		Description: "test",
		Completed:   false,
	}).Return(nil).Once()

	// test
	err := taskUsecase.UpdateTask(&entity.Task{
		ID:          1,
		Title:       "test",
		Description: "test",
		Completed:   false,
	})
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("DeleteTask", 1).Return(nil).Once()

	// test
	err := taskUsecase.DeleteTask(1)
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestCreateSubTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("CreateSubTask", 1, &entity.SubTask{
		Title: "test",
	}).Return(nil).Once()

	// test
	err := taskUsecase.CreateSubTask(1, &entity.SubTask{
		Title: "test",
	})
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestUpdateSubTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("UpdateSubTask", 1, &entity.SubTask{
		ID:    1,
		Title: "test",
	}).Return(nil).Once()

	// test
	err := taskUsecase.UpdateSubTask(1, &entity.SubTask{
		ID:    1,
		Title: "test",
	})
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestDeleteSubTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("DeleteSubTask", 1, 1).Return(nil).Once()

	// test
	err := taskUsecase.DeleteSubTask(1, 1)
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestCheckSubTask(t *testing.T) {
	// mock
	mockTaskRepo := mocks.TaskRepository{}
	cfg := &config.Config{}
	l := logger.New(cfg)

	// usecase
	taskUsecase := NewTaskUseCase(l, cfg, &mockTaskRepo)

	mockTaskRepo.On("CheckSubTask", 1, 1).Return(nil).Once()

	// test
	err := taskUsecase.CheckSubTask(1, 1)
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}
