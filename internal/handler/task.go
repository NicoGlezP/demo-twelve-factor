package handler

import (
	"demo-twelve/internal/request"
	"demo-twelve/internal/service"
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	service       *service.TaskService
	datadogClient *statsd.Client
}

func NewTaskHandler(taskService *service.TaskService, datadogClient *statsd.Client) *TaskHandler {
	return &TaskHandler{
		service:       taskService,
		datadogClient: datadogClient,
	}
}

func (t *TaskHandler) GetTasks(ctx *gin.Context) {
	tasks, err := t.service.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) AddTask(ctx *gin.Context) {
	var task *request.Task
	if err := ctx.ShouldBind(&task); err != nil {
		fmt.Printf(err.Error())
		// Contador de errores
		t.datadogClient.Incr("handler.errors", []string{"endpoint:AddTask"}, 1)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskCreated, err := t.service.CreateTask(*task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enviar la respuesta
	ctx.JSON(http.StatusCreated, taskCreated)
}

func (t *TaskHandler) ModifyTask(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented",
	})
}

func (t *TaskHandler) DeleteTask(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented",
	})
}
