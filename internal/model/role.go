package model

type Role string

const (
	RoleAdmin               Role = "admin"
	RoleStudent             Role = "student"
	RoleLecturer            Role = "lecturer"
	RoleLaboratoryAssistant Role = "laboratory_assistant"
)

type RoleModel struct {
	ID   int  `json:"id"`
	Name Role `json:"name"`
}

type UserRole struct {
	IDUser int `json:"id_user"`
	IDRole int `json:"id_role"`
}
