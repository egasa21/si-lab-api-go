package dto

import (
	"encoding/json"

	"github.com/google/uuid"
)

type CreatePracticumModuleContentRequest struct {
	IDModule int             `json:"id_module" validate:"required"`
	Title    string          `json:"title" validate:"required"`
	Content  json.RawMessage `json:"content" validate:"required"`
	Sequence int             `json:"sequence" validate:"required"`
}

type UpdatePracticumModuleContentRequest struct {
	IDModule   int       `json:"id_module"`
	Title      string    `json:"title"`
	Content    json.RawMessage `json:"content" validate:"required"`
	Sequence   int       `json:"sequence"`
	MaterialID uuid.UUID `json:"material_id"`
}


type PracticumModuleContentResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
