package dto

type StudentRegistrationRequest struct {
	StudentID   int   `json:"student_id" binding:"required"`
	PracticumIDs []int `json:"practicum_ids" binding:"required"`
}
