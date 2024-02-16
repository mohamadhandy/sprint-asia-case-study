package task

import (
	"service-task-list/config"
	"service-task-list/internal/entity"
	"service-task-list/internal/repository/mysql"
	"service-task-list/pkg/logger"
)

type ITaskUseCase interface {
	GetTaskList(completed int) ([]*entity.Task, error)
	CreateTask(task *entity.TaskRequest) error
	CheckTask(task *entity.Task) error
	UpdateTask(task *entity.Task) error
}

type TaskUseCase struct {
	l   *logger.Logger
	cfg *config.Config
	tr  mysql.TaskRepository
}

func NewTaskUseCase(l *logger.Logger, cfg *config.Config, tr mysql.TaskRepository) *TaskUseCase {
	return &TaskUseCase{l, cfg, tr}
}

func (t *TaskUseCase) GetTaskList(completed int) ([]*entity.Task, error) {
	tasks, err := t.tr.GetTaskList(completed)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskUseCase) CreateTask(task *entity.TaskRequest) error {
	err := t.tr.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUseCase) CheckTask(task *entity.Task) error {
	err := t.tr.CheckTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUseCase) UpdateTask(task *entity.Task) error {
	err := t.tr.UpdateTask(task)
	if err != nil {
		return err
	}
	return nil
}
