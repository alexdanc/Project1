package TaskService

import "time"

type Tasks struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Task      string     `json:"task"`
	IsDone    bool       `json:"isDone"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type PostTaskRequest struct {
	Task string `json:"task"`
}
