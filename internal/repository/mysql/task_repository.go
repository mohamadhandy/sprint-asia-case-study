package mysql

import (
	"database/sql"
	"service-task-list/config"
	"service-task-list/internal/entity"
	"service-task-list/pkg/logger"
)

type TaskRepository interface {
	GetTaskList(completed int) ([]*entity.Task, error)
	CreateTask(task *entity.TaskRequest) error
	CheckTask(id int) error
	UpdateTask(task *entity.Task) error
	DeleteTask(id int) error
}

type TaskMysqlRepo struct {
	l   *logger.Logger
	cfg *config.Config
	db  *sql.DB
}

func NewTaskMysqlRepo(l *logger.Logger, cfg *config.Config, db *sql.DB) *TaskMysqlRepo {
	return &TaskMysqlRepo{l, cfg, db}
}

func (t *TaskMysqlRepo) GetTaskList(completed int) ([]*entity.Task, error) {
	query := `SELECT id, title, description, created_at, completed FROM task WHERE completed = ? ORDER BY created_at DESC`
	rows, err := t.db.Query(query, completed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		task := new(entity.Task)
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TaskMysqlRepo) CreateTask(task *entity.TaskRequest) error {
	query := `INSERT INTO task (title, description, created_at, completed) VALUES (?, ?, ?, ?)`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Description, task.CreatedAt, task.Completed)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskMysqlRepo) CheckTask(id int) error {
	query := `UPDATE task SET completed = ? WHERE id = ?`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(true, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskMysqlRepo) UpdateTask(task *entity.Task) error {
	query := `UPDATE task SET title = ?, description = ? WHERE id = ?`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Description, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskMysqlRepo) DeleteTask(id int) error {
	query := `DELETE FROM task WHERE id = ?`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
