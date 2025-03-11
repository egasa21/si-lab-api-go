package model

import "time"

type UserPracticumCheckpoint struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	PracticumID int       `json:"practicum_id"`
	ModuleID    int       `json:"module_id"`
	ContentID   int       `json:"content_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
