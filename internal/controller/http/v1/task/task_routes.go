package task

import (
	"net/http"
	"service-task-list/config"
	"service-task-list/internal/usecase/task"
	"service-task-list/pkg/logger"

	"github.com/gorilla/mux"
)

type TaskRoutes struct {
	l   *logger.Logger
	cfg *config.Config
	tu  task.ITaskUseCase
}

func NewTaskRoutes(r *mux.Router, l *logger.Logger, cfg *config.Config, tu task.ITaskUseCase) {
	c := &TaskRoutes{l, cfg, tu}

	group := r.PathPrefix("/v1/tasks").Subrouter()
	group.HandleFunc("", c.GetTaskList).Methods(http.MethodGet)
	group.HandleFunc("/history", c.GetTaskListHistory).Methods(http.MethodGet)
	group.HandleFunc("/{task_id}/checklist", c.CheckTask).Methods(http.MethodPut)
	group.HandleFunc("/{task_id}", c.UpdateTask).Methods(http.MethodPut)
	group.HandleFunc("", c.CreateTask).Methods(http.MethodPost)
	group.HandleFunc("/{task_id}", c.DeleteTask).Methods(http.MethodDelete)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Test running Health function"))
	}).Methods(http.MethodGet)
}
