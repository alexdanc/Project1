package UserService

import (
	"context"

	"fmt"
)

type UserService interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	repoUser UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repoUser: r}
}

func (s *userService) GetAll(ctx context.Context) ([]User, error) {
	return s.repoUser.GetAllUsers(ctx)

}

func (s *userService) GetByID(ctx context.Context, id uint) (User, error) {
	return s.repoUser.GetUserByID(ctx, id)
}

func (s *userService) Create(ctx context.Context, user User) (User, error) {
	if user.Email == "" || user.Password == "" {
		return User{}, fmt.Errorf("email and password must not be empty")
	}
	return s.repoUser.CreateUser(ctx, user)
}

func (s *userService) Update(ctx context.Context, user User) error {
	return s.repoUser.UpdateUser(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.repoUser.DeleteUserByID(ctx, id)

}
