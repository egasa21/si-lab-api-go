package dto

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	IDStudent int    `json:"id_student"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
