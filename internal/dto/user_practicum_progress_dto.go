package dto

type CreateUserPracticumProgressRequest struct {
	UserID      int     `json:"user_id" validate:"required"`
	PracticumID int     `json:"practicum_id" validate:"required"`
	Progress    float64 `json:"progress" validate:"required,min=0,max=100"`
}

type UpdateUserPracticumProgressRequest struct {
	Progress float64 `json:"progress" validate:"required,min=0,max=100"`
}

type UserPracticumProgressResponse struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	PracticumID int     `json:"practicum_id"`
	Progress    float64 `json:"progress"`
	CompletedAt *string `json:"completed_at,omitempty"`
	LastUpdated string  `json:"last_updated"`
}
