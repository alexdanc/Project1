package handlers

import (
	"Project1/internal/TaskService"
	tasks "Project1/internal/Web/Tasks"
	"context"

	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestBodyHandlers struct {
	service TaskService.RequestBodyService
}

func (h *RequestBodyHandlers) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
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

func (h *RequestBodyHandlers) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body // теперь это *PostTaskRequest

	taskToCreate := TaskService.Tasks{
		Task:   taskRequest.Task,
		IsDone: false, // всегда false
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

func NewRequestBodyHandlers(s TaskService.RequestBodyService) *RequestBodyHandlers {
	return &RequestBodyHandlers{service: s}
}

func (h *RequestBodyHandlers) PatchHandler(c echo.Context) error {
	id := c.Param("id")
	var body TaskService.Tasks
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Ошибка парсинга JSON")
	}
	updatedTask, err := h.service.UpdateTask(id, body.Task)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func (h *RequestBodyHandlers) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTaskByID(id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
