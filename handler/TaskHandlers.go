package handlers

import (
	"Project1/internal/TaskService"
	tasks "Project1/internal/Web/Tasks"
	"context"
	"fmt"
	"strconv"
)

type TaskHandlers struct {
	service TaskService.TasksService
}

func NewTaskHandlers(s TaskService.TasksService) *TaskHandlers {
	return &TaskHandlers{service: s}
}

func (h *TaskHandlers) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandlers) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := TaskService.Tasks{
		Task:   *taskRequest.Task,
		IsDone: false,
	}

	createdTask, err := h.service.CreatesTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *TaskHandlers) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	idStr := strconv.Itoa(request.Id)

	err := h.service.DeleteTaskByID(idStr)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandlers) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	idStr := strconv.Itoa(request.Id)

	// Проверяем, что тело запроса не nil и поле Task не nil
	if request.Body == nil || request.Body.Task == nil {
		return nil, fmt.Errorf("missing 'task' field in request body")
	}

	updatedTask, err := h.service.UpdateTask(idStr, *request.Body.Task)
	if err != nil {
		return nil, err
	}

	// Формируем ответ в формате tasks.Task (OpenAPI)
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}
