//go:build wireinject
// +build wireinject

package main

import (
	"demo-twelve/internal/handler"
	"demo-twelve/internal/repository"
	"demo-twelve/internal/service"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Proveedor de la base de datos (DB)
func InitDB() (*gorm.DB, error) {
	dsn := "admin:SoyAdmin12@tcp(dev-demo.ceccfhe316xj.us-east-1.rds.amazonaws.com:3306)/tasks?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeTaskHandler() (*handler.TaskHandler, error) {
	wire.Build(
		InitDB,
		service.NewDatadogClient,
		repository.NewTaskRepository,
		service.NewTaskService,
		handler.NewTaskHandler,
	)
	return &handler.TaskHandler{}, nil
}
