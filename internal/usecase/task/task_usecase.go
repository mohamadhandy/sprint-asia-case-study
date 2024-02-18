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
	CheckTask(id int) error
	UpdateTask(task *entity.Task) error
	DeleteTask(id int) error
	CreateSubTask(taskId int, subTask *entity.SubTask) error
	UpdateSubTask(taskId int, subTask *entity.SubTask) error
	DeleteSubTask(taskId int, subTaskId int) error
	CheckSubTask(taskId int, subTaskId int) error
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

func (t *TaskUseCase) CheckTask(id int) error {
	err := t.tr.CheckTask(id)
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

func (t *TaskUseCase) DeleteTask(id int) error {
	err := t.tr.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUseCase) CreateSubTask(taskId int, subTask *entity.SubTask) error {
	err := t.tr.CreateSubTask(taskId, subTask)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUseCase) UpdateSubTask(taskId int, subTask *entity.SubTask) error {
	err := t.tr.UpdateSubTask(taskId, subTask)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUseCase) DeleteSubTask(taskId int, subTaskId int) error {
	err := t.tr.DeleteSubTask(taskId, subTaskId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUseCase) CheckSubTask(taskId int, subTaskId int) error {
	err := t.tr.CheckSubTask(taskId, subTaskId)
	if err != nil {
		return err
	}
	return nil
}
