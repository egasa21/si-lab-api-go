package dto

type StudentPracticumActivity struct {
	PracticumName   string `json:"practicum_name"`
	ModuleName      string `json:"module_name"`
	ModuleContentID int    `json:"module_content_id"`
}

type StudentCreateResponse struct {
	StudentID int `json:"student_id"`
}
