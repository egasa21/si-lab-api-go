package model

import (
	"encoding/json"
	"time"
	"github.com/google/uuid"
)

type PracticumModuleContent struct {
	IDContent int             `json:"id_content"`
	IDModule  int             `json:"id_module"`
	Title     string          `json:"title"`
	Content   json.RawMessage `json:"content"`
	Sequence  int             `json:"sequence"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	MaterialID uuid.UUID       `json:"material_id"`
}
