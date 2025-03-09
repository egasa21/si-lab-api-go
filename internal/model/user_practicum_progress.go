package model

import "time"

type UserPracticumProgress struct {
	ID          int        `json:"id"`
	UserID      int        `json:"id_user"`
	PracticumID int        `json:"id_practicum"`
	Progress    float64    `json:"progress"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	LastUpdated time.Time  `json:"last_updated_at"`
}
