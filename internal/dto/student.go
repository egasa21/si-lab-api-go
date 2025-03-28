package dto

type StudentPracticumActivity struct {
	ID                int    `json:"id"`
	PracticumName     string `json:"practicum_name"`
	ModuleName        string `json:"module_name"`
	ModuleContentName string `json:"module_content_name"`
	ModuleSequence    int    `json:"module_sequence"`
	ModuleContentID   int    `json:"module_content_id"`
}

type StudentCreateResponse struct {
	StudentID int `json:"student_id"`
}
