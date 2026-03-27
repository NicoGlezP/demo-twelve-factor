package service

import (
	"demo-twelve/internal/models"
	"demo-twelve/internal/repository"
	"demo-twelve/internal/request"
	"demo-twelve/internal/response"
	"demo-twelve/internal/utils"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) CreateTask(taskRequest request.Task) (*response.Task, error) {
	task := &models.Task{
		UUID:        utils.RandomString(UUIDLength),
		Name:        taskRequest.Name,
		Description: taskRequest.Description,
		Status:      statusActive,
	}

	createdTask, err := t.repo.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return &response.Task{
		UUID:        createdTask.UUID,
		Name:        createdTask.Name,
		Description: createdTask.Description,
		Status:      createdTask.Status,
	}, nil
}

func (t *TaskService) GetAllTasks() ([]response.Task, error) {
	tasks, err := t.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var taskResponses []response.Task

	for _, task := range tasks {
		taskResponses = append(taskResponses, response.Task{
			UUID:        task.UUID,
			Name:        task.Name,
			Description: task.Description,
			Status:      task.Status,
		})
	}

	return taskResponses, nil
}
