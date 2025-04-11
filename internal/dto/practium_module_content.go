package dto

import "encoding/json"

type CreatePracticumModuleContentRequest struct {
	IDModule int             `json:"id_module" validate:"required"`
	Title    string          `json:"title" validate:"required"`
	Content  json.RawMessage `json:"content" validate:"required"`
	Sequence int             `json:"sequence" validate:"required"`
}
