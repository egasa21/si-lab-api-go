package dto

type StudentPracticumActivity struct {
	ID                int    `json:"id"`
	PracticumName     string `json:"practicum_name"`
	PracticumProgress int    `json:"practicum_progress"`
	ModuleName        string `json:"module_name"`
	ModuleContentName string `json:"module_content_name"`
	ModuleSequence    int    `json:"module_sequence"`
	ModuleID          int    `json:"module_id"`
	ModuleContentID   int    `json:"module_content_id"`
}

type StudentSchedules struct {
	ID        int    `json:"id"`
	ClassName string `json:"class_name"`
	Day       string `json:"day"`
	ClassTime string `json:"class_time"`
}

type StudentCreateResponse struct {
	StudentID int `json:"student_id"`
}
