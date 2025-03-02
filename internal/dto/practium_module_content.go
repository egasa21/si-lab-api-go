package dto

type CreatePracticumModuleContentRequest struct {
	IDModule int    `json:"id_module" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Sequence int    `json:"sequence" validate:"required"`
}
