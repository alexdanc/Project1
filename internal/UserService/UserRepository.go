package UserService

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id uint) (User, error)
	CreateUser(ctx context.Context, user User) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUserByID(ctx context.Context, userID uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := u.DB.Find(&users).Error
	return users, err
}

func (u *userRepository) GetUserByID(ctx context.Context, id uint) (User, error) {
	var user User
	err := u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) CreateUser(ctx context.Context, user User) (User, error) {
	// Проверка на уникальность email
	var existingUser User
	if err := r.DB.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return User{}, fmt.Errorf("user with this email already exists")
	}

	// Если пользователь не найден — создаем
	if err := r.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user User) error {
	err := u.DB.Save(&user).Error
	return err
}

func (u *userRepository) DeleteUserByID(ctx context.Context, id uint) error {
	err := u.DB.Delete(&User{}, id).Error
	return err
}
