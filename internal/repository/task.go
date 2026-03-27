package repository

import (
	"demo-twelve/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	result := t.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (t *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := t.db.Find(&tasks)

	return tasks, result.Error
}

func (t *TaskRepository) ChangeToDone(task *models.Task) (*models.Task, error) {
	// TODO: Implementar para fines de la clase
	return nil, nil
}
