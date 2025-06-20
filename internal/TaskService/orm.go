package TaskService

type RequestBody struct {
	ID     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
