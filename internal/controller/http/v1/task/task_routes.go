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

	group := r.PathPrefix("/v1").Subrouter()
	group.HandleFunc("/tasks", c.GetTaskList).Methods(http.MethodGet)
	group.HandleFunc("/tasks/history", c.GetTaskListHistory).Methods(http.MethodGet)
	group.HandleFunc("/tasks/checklist", c.CheckTask).Methods(http.MethodPut)
	group.HandleFunc("/tasks", c.UpdateTask).Methods(http.MethodPut)
	group.HandleFunc("/tasks", c.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Test running Health function"))
	}).Methods(http.MethodGet)
}
