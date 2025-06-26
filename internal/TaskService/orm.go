package TaskService

import (
	"time"
)

type Tasks struct {
	ID         uint      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Task       string    `json:"task"`
	IsDone     bool      `json:"isDone"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}
