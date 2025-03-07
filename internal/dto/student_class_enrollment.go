package dto

// StudentEnrollmentRequest represents the request payload for enrolling a student in a class
type StudentEnrollmentRequest struct {
	StudentID int `json:"student_id"`
	ClassID   int `json:"class_id"`
}

// StudentEnrollmentResponse represents the response for a student class enrollment
type StudentEnrollmentResponse struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	ClassID   int `json:"class_id"`
}
