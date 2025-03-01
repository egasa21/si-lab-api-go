package model

import "time"

type PracticumModule struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PracticumID int       `json:"practicum_id"`
}