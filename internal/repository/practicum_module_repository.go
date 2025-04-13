package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type PracticumModuleRepository interface {
	CreateModule(module *model.PracticumModule) (*model.PracticumModule, error)
	GetModuleByID(id int) (*model.PracticumModule, error)
	GetModuleByIDs(ids []int) ([]model.PracticumModule, error)
	GetModulesByPracticumID(practicumID, page, limit int) ([]model.PracticumModule, int, error)
}

type practicumModuleRepository struct {
	db *sql.DB
}

func NewPracticumModuleRepository(db *sql.DB) PracticumModuleRepository {
	return &practicumModuleRepository{db: db}
}

func (r *practicumModuleRepository) CreateModule(module *model.PracticumModule) (*model.PracticumModule, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
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
        INSERT INTO practicum_modules (title, practicum_id)
        VALUES ($1, $2)
        RETURNING id
    `
	err = tx.QueryRow(query, module.Title, module.PracticumID).Scan(&module.ID)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (r *practicumModuleRepository) GetModuleByID(id int) (*model.PracticumModule, error) {
	var module model.PracticumModule
	err := r.db.QueryRow(
		"SELECT id, title, practicum_id, created_at, updated_at FROM practicum_modules WHERE id = $1", id,
	).Scan(&module.ID, &module.Title, &module.PracticumID, &module.CreatedAt, &module.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *practicumModuleRepository) GetModulesByPracticumID(practicumID, page, limit int) ([]model.PracticumModule, int, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Query(
		"SELECT id, title, practicum_id, created_at, updated_at FROM practicum_modules WHERE practicum_id = $1 LIMIT $2 OFFSET $3",
		practicumID, limit, offset,
	)
	if err != nil {
		log.Error().Err(err)
		return nil, 0, err
	}
	defer rows.Close()

	var modules []model.PracticumModule
	for rows.Next() {
		var module model.PracticumModule
		if err := rows.Scan(&module.ID, &module.Title, &module.PracticumID, &module.CreatedAt, &module.UpdatedAt); err != nil {
			return nil, 0, err
		}
		modules = append(modules, module)
	}

	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM practicum_modules WHERE practicum_id = $1", practicumID).Scan(&total)
	if err != nil {
		log.Error().Err(err)
		return nil, 0, err
	}

	return modules, total, nil
}

func (r *practicumModuleRepository) GetModuleByIDs(ids []int) ([]model.PracticumModule, error) {
	if len(ids) == 0 {
		return []model.PracticumModule{}, nil
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	inClause := strings.Join(placeholders, ",")
	query := fmt.Sprintf("SELECT id, title, created_at, updated_at FROM practicum_modules WHERE id IN (%s)", inClause)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	practicumModules := []model.PracticumModule{}
	for rows.Next() {
		practicumModule := model.PracticumModule{}
		if err := rows.Scan(&practicumModule.ID, &practicumModule.Title, &practicumModule.CreatedAt, &practicumModule.UpdatedAt); err != nil {
			return nil, err
		}
		practicumModules = append(practicumModules, practicumModule)
	}

	return practicumModules, nil
}
