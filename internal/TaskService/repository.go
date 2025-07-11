package TaskService

import "gorm.io/gorm"

type RequestBodyRepository interface {
	CreateTask(req *Tasks) error
	GetAllTasks() ([]Tasks, error)
	GetTaskByID(id string) (Tasks, error)
	UpdateTask(req Tasks) error
	DeleteTaskByID(id string) error
}

type TaskRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) RequestBodyRepository {
	return &TaskRepository{DB: DB}
}

func (r *TaskRepository) CreateTask(req *Tasks) error {
	result := r.DB.Create(req) // передаем указатель, чтобы GORM мог заполнить поля
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TaskRepository) GetAllTasks() ([]Tasks, error) {
	var task []Tasks
	err := r.DB.Find(&task).Error
	return task, err
}

func (r *TaskRepository) GetTaskByID(id string) (Tasks, error) {
	var tas Tasks
	err := r.DB.First(&tas, id).Error
	return tas, err

}

func (r *TaskRepository) UpdateTask(req Tasks) error {
	return r.DB.Save(&req).Error
}

func (r *TaskRepository) DeleteTaskByID(id string) error {
	return r.DB.Delete(&Tasks{}, "id = ?", id).Error
}
