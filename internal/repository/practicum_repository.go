package repository

import (
	"database/sql"

	"github.com/egasa21/si-lab-api-go/internal/model"
)

type PracticumRepository interface {
	CreatePracticum(practicum *model.Practicum) error
}

type practicumRepository struct {
	db *sql.DB
}

func NewPracticumRepository(db *sql.DB) PracticumRepository {
	return &practicumRepository{db: db}
}

func (r *practicumRepository) CreatePracticum(practicum *model.Practicum) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
        INSERT INTO practicums (name, code, description, credits, semester)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id_practicum
    `

	_, err = tx.Exec(query, practicum.Name, practicum.Code, practicum.Description, practicum.Credits, practicum.Semester)
	if err != nil {
		return err
	}

	return nil
}
