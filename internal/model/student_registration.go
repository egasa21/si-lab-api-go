package model

import "time"

type StudentRegistration struct {
	IDStudentRegistration int       `json:"id_student_registration"`
	StudentID             int       `json:"student_id"`
	PracticumID           int       `json:"practicum_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
