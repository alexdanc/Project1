package TaskService

import "time"

type RequestBodyService interface {
	CreatesTask(task Tasks) (Tasks, error)
	GetAllTasks() ([]Tasks, error)
	GetTaskByID(id string) (Tasks, error)
	UpdateTask(id string, body string) (Tasks, error)
	DeleteTaskByID(id string) error
}

type TaskService struct {
	repo RequestBodyRepository
}

func NewTaskService(r RequestBodyRepository) RequestBodyService {
	return &TaskService{repo: r}
}

func (s *TaskService) CreatesTask(task Tasks) (Tasks, error) {
	newTask := Tasks{
		Task:       task.Task,
		IsDone:     false,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Deleted_at: time.Now(),
	}

	err := s.repo.CreateTask(newTask)
	if err != nil {
		return Tasks{}, err
	}

	return newTask, nil
}

func (s *TaskService) GetAllTasks() ([]Tasks, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id string) (Tasks, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(id, body string) (Tasks, error) {
	tas, err := s.repo.GetTaskByID(id)
	if err != nil {
		return tas, err
	}
	tas.Task = body
	if err := s.repo.UpdateTask(tas); err != nil {
		return tas, err
	}
	return tas, nil
}

func (s *TaskService) DeleteTaskByID(id string) error {
	return s.repo.DeleteTaskByID(id)
}
