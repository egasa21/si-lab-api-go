package model

import "time"

type Practicum struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Credits     string    `json:"credits"`
	Semester    string    `json:"semester"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Material struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type ModuleWithMaterials struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Materials []Material `json:"materials"`
}

type PracticumWithMaterial struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Code        string                 `json:"code"`
	Description string                 `json:"description"`
	Credits     string                 `json:"credits"`
	Semester    string                 `json:"semester"`
	Modules     []*ModuleWithMaterials `json:"modules"` 
}

