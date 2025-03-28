package repository

import (
	"database/sql"

	"github.com/egasa21/si-lab-api-go/internal/model"
)

type StudentRepository interface {
	GetAllStudents(page, limit int) ([]model.Student, int, error)
	GetStudentByID(id int) (*model.Student, error)
	CreateStudent(student *model.Student) (int, error)
	GetStudentByStudentID(student_id_number string) (*model.Student, error)
}

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) GetAllStudents(page, limit int) ([]model.Student, int, error) {
	// Calculate the offset based on the page and limit
	offset := (page - 1) * limit

	// Query to get the students with pagination
	rows, err := r.db.Query(
		"SELECT id, student_id_number, name, study_plan_file, created_at, updated_at FROM students LIMIT $1 OFFSET $2",
		limit, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var student model.Student
		if err := rows.Scan(&student.ID, &student.StudentIDNumber, &student.Name, &student.StudyPlanFile, &student.CreatedAt, &student.UpdatedAt); err != nil {
			return nil, 0, err
		}
		students = append(students, student)
	}

	// Get the total number of students for pagination
	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM students").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

func (r *studentRepository) GetStudentByID(id int) (*model.Student, error) {
	var student model.Student
	err := r.db.QueryRow("SELECT id, student_id_number, name, study_plan_file, created_at, updated_at FROM students WHERE id = $1", id).
		Scan(&student.ID, &student.StudentIDNumber, &student.Name, &student.StudyPlanFile, &student.CreatedAt, &student.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) GetStudentByStudentID(student_id_number string) (*model.Student, error) {
	var student model.Student
	err := r.db.QueryRow("SELECT id, student_id_number, name, study_plan_file, created_at, updated_at FROM students WHERE student_id_number = $1", student_id_number).
		Scan(&student.ID, &student.StudentIDNumber, &student.Name, &student.StudyPlanFile, &student.CreatedAt, &student.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil // No student found, return nil without error
	} else if err != nil {
		return nil, err // Some other error occurred
	}
	return &student, nil // Student found
}

func (r *studentRepository) CreateStudent(student *model.Student) (int, error) {
	// start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	err = tx.QueryRow("INSERT INTO students (student_id_number, name, study_plan_file, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		student.StudentIDNumber, student.Name, student.StudyPlanFile, student.CreatedAt).Scan(&id)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}
