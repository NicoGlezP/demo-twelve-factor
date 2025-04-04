package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	_apiVersion = "/api/v1"
)

const (
	_tasksPath = _apiVersion + "/tasks"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	taskHandler, err := InitializeTaskHandler()
	if err != nil {
		panic(err)
	}

	router.GET(_tasksPath, taskHandler.GetTasks)
	router.POST(_tasksPath, taskHandler.AddTask)
	router.PUT(_tasksPath, taskHandler.ModifyTask)
	router.DELETE(_tasksPath, taskHandler.DeleteTask)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
