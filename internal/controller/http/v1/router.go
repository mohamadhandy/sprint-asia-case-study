package v1

import (
	"service-task-list/config"
	task_http "service-task-list/internal/controller/http/v1/task"
	"service-task-list/internal/usecase/task"
	"service-task-list/pkg/logger"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, l *logger.Logger, cfg *config.Config, tu task.ITaskUseCase) {
	{
		task_http.NewTaskRoutes(r, l, cfg, tu)
	}
}
