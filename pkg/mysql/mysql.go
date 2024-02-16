package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"service-task-list/config"
	"service-task-list/pkg/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(driverName string, cfg *config.Config, l *logger.Logger) *sql.DB {
	db, err := sql.Open(driverName, cfg.MYSQL.URL)
	if err != nil {
		log.Fatalf(fmt.Sprintf("couldn't connect to database connection: %v", err))
	}

	db.SetMaxIdleConns(cfg.MYSQL.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MYSQL.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MYSQL.MaxLifetimeConns) * time.Second)

	return db
}
