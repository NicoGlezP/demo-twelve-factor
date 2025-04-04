package handler

import (
	"demo-twelve/internal/request"
	"demo-twelve/internal/service"
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	// Iniciar el temporizador para medir la latencia
	start := time.Now()
	defer func() {
		// Calcular y registrar el tiempo de latencia
		t.datadogClient.Timing("handler.latency", time.Since(start), []string{"endpoint:GetTasks"}, 1)
	}()

	// Contador de solicitudes
	t.datadogClient.Incr("handler.requests", []string{"endpoint:GetTasks"}, 1)

	// Obtener las tareas
	tasks, err := t.service.GetAllTasks()
	if err != nil {
		// Contador de errores
		t.datadogClient.Incr("handler.errors", []string{"endpoint:GetTasks"}, 1)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registrar el tamaño de la respuesta
	t.datadogClient.Gauge("handler.response_size", float64(len(tasks)), []string{"endpoint:GetTasks"}, 1)

	// Enviar la respuesta
	ctx.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) AddTask(ctx *gin.Context) {
	// Iniciar el temporizador para medir la latencia
	start := time.Now()
	defer func() {
		// Calcular y registrar el tiempo de latencia
		t.datadogClient.Timing("handler.latency", time.Since(start), []string{"endpoint:AddTask"}, 1)
	}()

	// Contador de solicitudes
	t.datadogClient.Incr("handler.requests", []string{"endpoint:AddTask"}, 1)

	// Bind de datos de la solicitud
	var task *request.Task
	if err := ctx.ShouldBind(&task); err != nil {
		fmt.Printf(err.Error())
		// Contador de errores
		t.datadogClient.Incr("handler.errors", []string{"endpoint:AddTask"}, 1)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la tarea
	supplierCreated, err := t.service.CreateTask(*task)
	if err != nil {
		// Contador de errores
		t.datadogClient.Incr("handler.errors", []string{"endpoint:AddTask"}, 1)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Registrar el tamaño de la respuesta
	t.datadogClient.Gauge("handler.response_size", float64(len(fmt.Sprintf("%v", supplierCreated))), []string{"endpoint:AddTask"}, 1)

	// Enviar la respuesta
	ctx.JSON(http.StatusCreated, supplierCreated)
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
