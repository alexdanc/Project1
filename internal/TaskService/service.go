package TaskService

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
	newTask := &Tasks{
		Task:   task.Task,
		IsDone: false, // Можно либо всегда ставить false, либо брать из task.IsDone
	}

	err := s.repo.CreateTask(newTask) // передаем указатель
	if err != nil {
		return Tasks{}, err
	}

	return *newTask, nil // возвращаем обновленную структуру с ID
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
