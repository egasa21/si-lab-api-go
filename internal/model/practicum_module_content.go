package model

import "time"

type PracticumModuleContent struct {
	IDContent int       `json:"id_content"`
	IDModule  int       `json:"id_module"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Sequence  int       `json:"sequence"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
