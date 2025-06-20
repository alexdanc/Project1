package TaskService

import "time"

type RequestBody struct {
	ID         uint      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Task       string    `json:"task"`
	IsDone     bool      `json:"isDone"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
	deleted_at time.Time `sql:"deleted_at"`
}
