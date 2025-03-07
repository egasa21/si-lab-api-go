package model

import "time"

type StudentClassEnrollment struct {
	ID        int       `json:"id"`
	ClassID   int       `json:"class_id"`
	StudentID int       `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
