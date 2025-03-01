package dto

type CreatePracticumModuleRequest struct {
	Title       string `json:"title" binding:"required"`
	PracticumID int    `json:"practicum_id" binding:"required"`
}
