package TaskService

import "gorm.io/gorm"

type RequestBodyRepository interface {
	CreateTask(req Tasks) error
	GetAllTasks() ([]Tasks, error)
	GetTaskByID(id string) (Tasks, error)
	UpdateTask(req Tasks) error
	DeleteTaskByID(id string) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RequestBodyRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(req Tasks) error {
	result := r.db.Create(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TaskRepository) GetAllTasks() ([]Tasks, error) {
	var tasks []Tasks
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id string) (Tasks, error) {
	var tas Tasks
	err := r.db.First(&tas, id).Error
	return tas, err

}

func (r *TaskRepository) UpdateTask(req Tasks) error {
	return r.db.Save(&req).Error
}

func (r *TaskRepository) DeleteTaskByID(id string) error {
	return r.db.Delete(&Tasks{}, "id = ?", id).Error
}
