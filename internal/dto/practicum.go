package dto

type MaterialResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type ModuleResponse struct {
	ID        uint               `json:"id"`
	Title     string             `json:"title"`
	Materials []MaterialResponse `json:"materials"`
}

type PracticumWithMaterialResponse struct {
	ID          uint             `json:"id"`
	Code        string           `json:"code"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Credits     string           `json:"credits"`
	Semester    string           `json:"semester"`
	Modules     []ModuleResponse `json:"modules"`
}
