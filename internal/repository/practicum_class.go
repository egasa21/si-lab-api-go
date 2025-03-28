package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type PracticumClassRepository interface {
	CreateClass(class *model.PracticumClass) error
	GetClassByID(id int) (*model.PracticumClass, error)
	GetClassByIDs(ids []int) ([]model.PracticumClass, error)
	GetClassesByPracticumID(practicumID int) ([]model.PracticumClass, error)
	UpdateClass(class *model.PracticumClass) error
	DeleteClass(id int) error
}

type practicumClassRepository struct {
	db *sql.DB
}

func NewPracticumClassRepository(db *sql.DB) PracticumClassRepository {
	return &practicumClassRepository{db: db}
}

func (r *practicumClassRepository) CreateClass(class *model.PracticumClass) error {
	query := `
		INSERT INTO practicum_class (practicum_id, name, quota, day, time, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id_practicum_class
	`
	err := r.db.QueryRow(query, class.PracticumID, class.Name, class.Quota, class.Day, class.Time).
		Scan(&class.IDPracticumClass)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create practicum class")
		return err
	}
	return nil
}

func (r *practicumClassRepository) GetClassByID(id int) (*model.PracticumClass, error) {
	var class model.PracticumClass
	query := `
		SELECT id_practicum_class, practicum_id, name, quota, day, time, created_at, updated_at 
		FROM practicum_class 
		WHERE id_practicum_class = $1
	`
	err := r.db.QueryRow(query, id).
		Scan(&class.IDPracticumClass, &class.PracticumID, &class.Name, &class.Quota, &class.Day, &class.Time, &class.CreatedAt, &class.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get practicum class by ID")
		return nil, err
	}
	return &class, nil
}

func (r *practicumClassRepository) GetClassesByPracticumID(practicumID int) ([]model.PracticumClass, error) {
	query := `
		SELECT id_practicum_class, practicum_id, name, quota, day, time, created_at, updated_at 
		FROM practicum_class 
		WHERE practicum_id = $1
	`
	rows, err := r.db.Query(query, practicumID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get practicum classes by practicum ID")
		return nil, err
	}
	defer rows.Close()

	var classes []model.PracticumClass
	for rows.Next() {
		var class model.PracticumClass
		if err := rows.Scan(&class.IDPracticumClass, &class.PracticumID, &class.Name, &class.Quota, &class.Day, &class.Time, &class.CreatedAt, &class.UpdatedAt); err != nil {
			log.Error().Err(err).Msg("Failed to scan practicum class")
			return nil, err
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func (r *practicumClassRepository) UpdateClass(class *model.PracticumClass) error {
	query := `
		UPDATE practicum_class 
		SET practicum_id = $1, name = $2, quota = $3, day = $4, time = $5, updated_at = NOW() 
		WHERE id_practicum_class = $6
	`
	_, err := r.db.Exec(query, class.PracticumID, class.Name, class.Quota, class.Day, class.Time, class.IDPracticumClass)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update practicum class")
		return err
	}
	return nil
}

func (r *practicumClassRepository) DeleteClass(id int) error {
	query := `DELETE FROM practicum_class WHERE id_practicum_class = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete practicum class")
		return err
	}
	return nil
}

func (r *practicumClassRepository) GetClassByIDs(ids []int) ([]model.PracticumClass, error) {
	if len(ids) == 0 {
		return []model.PracticumClass{}, nil
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	inClause := strings.Join(placeholders, ",")
	query := fmt.Sprintf("SELECT id_practicum_class, practicum_id, name, quota, day, time, created_at, updated_at FROM practicum_class WHERE id_practicum_class IN (%s)", inClause)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	practicumClasses := []model.PracticumClass{}
	for rows.Next() {
		practicumClass := model.PracticumClass{}
		if err := rows.Scan(&practicumClass.IDPracticumClass, &practicumClass.PracticumID, &practicumClass.Name, &practicumClass.Quota, &practicumClass.Day, &practicumClass.Time, &practicumClass.CreatedAt, &practicumClass.UpdatedAt); err != nil {
			return nil, err
		}

		practicumClasses = append(practicumClasses, practicumClass)
	}

	return practicumClasses, nil
}
