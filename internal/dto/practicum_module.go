package dto

type CreatePracticumModuleRequest struct {
	Title       string `json:"title" binding:"required"`
	PracticumID int    `json:"practicum_id" binding:"required"`
}

type PracticumModuleResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
