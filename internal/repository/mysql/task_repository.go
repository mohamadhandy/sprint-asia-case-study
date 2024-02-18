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
	CreateSubTask(taskId int, subTask *entity.SubTask) error
	UpdateSubTask(taskId int, subTask *entity.SubTask) error
	DeleteSubTask(taskId int, subTaskId int) error
	GetAllSubTaskByCompleted(taskId int, completed string) ([]*entity.SubTask, error)
	CheckSubTask(taskId int, subTaskId int) error
	GetPercentageSubTask(taskId int) (int, error)
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

func (t *TaskMysqlRepo) CreateSubTask(taskId int, subTask *entity.SubTask) error {
	query := `INSERT INTO sub_task (task_id, title, completed) VALUES (?, ?, ?)`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskId, subTask.Title, false)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskMysqlRepo) UpdateSubTask(taskId int, subTask *entity.SubTask) error {
	query := `UPDATE sub_task SET title = ?, completed = ? WHERE id = ? AND task_id = ?`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subTask.Title, false, subTask.ID, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskMysqlRepo) DeleteSubTask(taskId int, subTaskId int) error {
	query := `DELETE FROM sub_task WHERE id = ? AND task_id = ?`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subTaskId, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskMysqlRepo) GetAllSubTaskByCompleted(taskId int, completed string) ([]*entity.SubTask, error) {
	var query string
	var rows *sql.Rows
	var err error
	if completed == "true" {
		query = `SELECT id, title, completed FROM sub_task WHERE task_id = ? AND completed = ?`
		rows, err = t.db.Query(query, taskId, true)
	} else if completed == "false" {
		query = `SELECT id, title, completed FROM sub_task WHERE task_id = ? AND completed = ?`
		rows, err = t.db.Query(query, taskId, false)
	} else {
		query = `SELECT id, title, completed FROM sub_task WHERE task_id = ?`
		rows, err = t.db.Query(query, taskId)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subTasks []*entity.SubTask
	for rows.Next() {
		subTask := new(entity.SubTask)
		err := rows.Scan(&subTask.ID, &subTask.Title, &subTask.Completed)
		if err != nil {
			return nil, err
		}
		subTasks = append(subTasks, subTask)
	}
	return subTasks, nil
}

func (t *TaskMysqlRepo) CheckSubTask(taskId int, subTaskId int) error {
	query := `UPDATE sub_task SET completed = ? WHERE id = ? AND task_id = ?`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(true, subTaskId, taskId)
	if err != nil {
		return err
	}
	// check subtask is all completed
	subTaskCompleted, err := t.GetAllSubTaskByCompleted(taskId, "true")
	if err != nil {
		return err
	}
	allSubTask, err := t.GetAllSubTaskByCompleted(taskId, "all")
	if err != nil {
		return err
	} else if len(subTaskCompleted) == len(allSubTask) {
		err = t.CheckTask(taskId)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (t *TaskMysqlRepo) GetPercentageSubTask(taskId int) (int, error) {
	subTaskCompleted, err := t.GetAllSubTaskByCompleted(taskId, "true")
	if err != nil {
		return 0, err
	}
	allSubTask, err := t.GetAllSubTaskByCompleted(taskId, "all")
	if err != nil {
		return 0, err
	}
	percentage := (len(subTaskCompleted) * 100) / len(allSubTask)
	return percentage, nil
}
