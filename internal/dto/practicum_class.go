package dto

type CreatePracticumClassRequest struct {
	PracticumID int    `json:"practicum_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Quota       int    `json:"quota" validate:"required,min=1"`
	Day         string `json:"day" validate:"required"`
	Time        string `json:"time" validate:"required"`
}

type UpdatePracticumClassRequest struct {
	PracticumID      int    `json:"practicum_id"`
	Name             string `json:"name"`
	Quota            int    `json:"quota"`
	Day              string `json:"day"`
	Time             string `json:"time"`
}
