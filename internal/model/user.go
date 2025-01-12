package model

import "time"

type User struct {
	IDUser    int         `json:"id_user"`
	Email     string      `json:"email"`
	Password  string      `json:"-"`
	IDStudent int         `json:"id_student"`
	Roles     []RoleModel `json:"roles"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
