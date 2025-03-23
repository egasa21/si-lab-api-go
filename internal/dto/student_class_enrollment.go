package dto

type StudentEnrollmentRequest struct {
	StudentID int `json:"student_id"`
	ClassID   int `json:"class_id"`
}

type StudentEnrollmentResponse struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	ClassID   int `json:"class_id"`
}

// type StudentEnrollmentResponseTest struct {
// 	ID        int `json:"id"`
// 	StudentID int `json:"student_id"`
// 	ClassID   int `json:"class_id"`
// }