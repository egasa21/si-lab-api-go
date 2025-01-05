package model

import "time"

type Student struct {
	ID              int       `json:"id"`
	StudentIDNumber string    `json:"student_id_number"`
	Name            string    `json:"name"`
	StudyPlanFile   string    `json:"study_plan_file"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
