package dto

type CreateUserPracticumCheckpointRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	PracticumID int `json:"practicum_id" binding:"required"`
	ModuleID  int `json:"module_id" binding:"required"`
	ContentID int `json:"content_id" binding:"required"`
}
