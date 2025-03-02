package model

import "time"

type PracticumClass struct {
	IDPracticumClass int       `json:"id_practicum_class"`
	PracticumID      int       `json:"practicum_id"`
	Name             string    `json:"name"`
	Quota            int       `json:"quota"`
	Day              string    `json:"day"`
	Time             string    `json:"time"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
