package TaskService

type TasksService interface {
	CreatesTask(task Tasks) (Tasks, error)
	GetAllTasks() ([]Tasks, error)
	GetTaskByID(id uint) (Tasks, error)
	UpdateTask(id uint, body string) (Tasks, error)
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userId uint) ([]Tasks, error)
}

type TaskService struct {
	repoTask TasksRepository
}

func NewTaskService(r TasksRepository) TasksService {
	return &TaskService{repoTask: r}
}

func (s *TaskService) GetAllTasks() ([]Tasks, error) {
	return s.repoTask.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id uint) (Tasks, error) {
	return s.repoTask.GetTaskByID(id)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]Tasks, error) {
	return s.repoTask.GetByUserID(userID)
}

func (s *TaskService) CreatesTask(task Tasks) (Tasks, error) {
	newTask := &Tasks{
		Task:   task.Task,
		IsDone: false,
		UserID: task.UserID,
	}

	err := s.repoTask.CreateTask(newTask)
	if err != nil {
		return Tasks{}, err
	}

	return *newTask, nil
}

func (s *TaskService) UpdateTask(id uint, body string) (Tasks, error) {
	tas, err := s.repoTask.GetTaskByID(id)
	if err != nil {
		return tas, err
	}
	tas.Task = body
	if err := s.repoTask.UpdateTask(tas); err != nil {
		return tas, err
	}
	return tas, nil
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repoTask.DeleteTaskByID(id)
}
