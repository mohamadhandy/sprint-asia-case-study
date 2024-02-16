package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"service-task-list/config"
	v1 "service-task-list/internal/controller/http/v1"
	mysql_repository "service-task-list/internal/repository/mysql"
	"service-task-list/internal/usecase/task"
	"service-task-list/pkg/httpserver"
	"service-task-list/pkg/logger"
	"service-task-list/pkg/mysql"
	"syscall"

	"github.com/gorilla/mux"
)

func Run(cfg *config.Config) {
	fmt.Println("Running Service-TaskList")

	var err error
	l := logger.New(cfg)

	// Mysql
	db := mysql.New(cfg.MYSQL.MysqlDriverName, cfg, l)
	defer db.Close()

	// Repository
	// consumerRepository := mysql_repository.NewConsumerMysqlRepository(l, cfg, db)
	taskRepository := mysql_repository.NewTaskMysqlRepo(l, cfg, db)

	// Usecase
	// consumerUsecase := consumer.NewConsumerUsecase(l, cfg, consumerRepository)
	taskUsecase := task.NewTaskUseCase(l, cfg, taskRepository)

	// HTTP Server
	handler := mux.NewRouter()
	v1.NewRouter(handler, l, cfg, taskUsecase)
	httpServer := httpserver.New(handler, cfg, httpserver.Port(cfg.HTTPServer.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
