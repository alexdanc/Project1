package UserService

import (
	"Project1/internal/TaskService"
	"time"
)

type User struct {
	ID        uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string              `json:"email"`
	Password  string              `json:"password"`
	CreatedAt time.Time           `json:"createdAt,omitempty"`
	UpdatedAt time.Time           `json:"updatedAt,omitempty"`
	DeletedAt *time.Time          `json:"deletedAt,omitempty"`
	Tasks     []TaskService.Tasks `gorm:"foreignKey:UserID"`
}
