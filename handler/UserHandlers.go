package handlers

import (
	"Project1/internal/UserService"
	users "Project1/internal/Web/Users"
	"context"
	"fmt"
)

type UserHandlers struct {
	service UserService.UserService
}

func NewUserHandler(s UserService.UserService) *UserHandlers {
	return &UserHandlers{service: s}
}

func (h *UserHandlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, u := range allUsers {
		user := users.User{
			Id:       int(u.ID),
			Email:    u.Email,
			Password: u.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	if userRequest.Email == "" || userRequest.Password == "" {
		return nil, fmt.Errorf("email and password cannot be empty")
	}

	userToCreate := UserService.User{
		Password: userRequest.Password,
		Email:    userRequest.Email,
	}

	createdUser, err := h.service.Create(ctx, userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       int(createdUser.ID),
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}
	return response, nil
}

func (h *UserHandlers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := uint(request.Id)

	if request.Body == nil {
		return nil, fmt.Errorf("missing request body")
	}

	// Получаем пользователя
	user, err := h.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Обновляем поля, если они пришли
	if request.Body.Email != nil {
		user.Email = *request.Body.Email
	}
	if request.Body.Password != nil {
		user.Password = *request.Body.Password
	}

	// Сохраняем изменения
	if err := h.service.Update(ctx, user); err != nil {
		return nil, err
	}

	// Формируем ответ
	response := users.PatchUsersId200JSONResponse{
		Id:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}
	return response, nil
}

func (h *UserHandlers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := uint(request.Id)

	err := h.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
