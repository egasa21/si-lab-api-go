package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type PracticumModuleContentRepository interface {
	CreateContent(content *model.PracticumModuleContent) error
	GetContentByID(id int) (*model.PracticumModuleContent, error)
	GetContentByIDs(ids []int) ([]model.PracticumModuleContent, error)
	GetContentsByModuleID(moduleID, page, limit int) ([]model.PracticumModuleContent, int, error)
}

type practicumModuleContentRepository struct {
	db *sql.DB
}

func NewPracticumModuleContentRepository(db *sql.DB) PracticumModuleContentRepository {
	return &practicumModuleContentRepository{db: db}
}

func (r *practicumModuleContentRepository) CreateContent(content *model.PracticumModuleContent) error {
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
		INSERT INTO practicum_module_content (id_module, title, content, sequence)
		VALUES ($1, $2, $3, $4)
		RETURNING id_content
	`
	_, err = tx.Exec(query, content.IDModule, content.Title, content.Content, content.Sequence)
	if err != nil {
		return err
	}

	return nil
}

func (r *practicumModuleContentRepository) GetContentByID(id int) (*model.PracticumModuleContent, error) {
	var content model.PracticumModuleContent
	err := r.db.QueryRow(
		"SELECT id_content, id_module, title, content, sequence, created_at, updated_at FROM practicum_module_content WHERE id_content = $1", id,
	).Scan(&content.IDContent, &content.IDModule, &content.Title, &content.Content, &content.Sequence, &content.CreatedAt, &content.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *practicumModuleContentRepository) GetContentsByModuleID(moduleID, page, limit int) ([]model.PracticumModuleContent, int, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Query(
		"SELECT id_content, id_module, title, content, sequence, created_at, updated_at FROM practicum_module_content WHERE id_module = $1 ORDER BY sequence LIMIT $2 OFFSET $3",
		moduleID, limit, offset,
	)
	if err != nil {
		log.Error().Err(err)
		return nil, 0, err
	}
	defer rows.Close()

	var contents []model.PracticumModuleContent
	for rows.Next() {
		var content model.PracticumModuleContent
		if err := rows.Scan(&content.IDContent, &content.IDModule, &content.Title, &content.Content, &content.Sequence, &content.CreatedAt, &content.UpdatedAt); err != nil {
			return nil, 0, err
		}
		contents = append(contents, content)
	}

	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM practicum_module_content WHERE id_module = $1", moduleID).Scan(&total)
	if err != nil {
		log.Error().Err(err)
		return nil, 0, err
	}

	return contents, total, nil
}

func (r *practicumModuleContentRepository) GetContentByIDs(ids []int) ([]model.PracticumModuleContent, error) {
	if len(ids) == 0 {
		return []model.PracticumModuleContent{}, nil
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	inClause := strings.Join(placeholders, ",")
	query := fmt.Sprintf("SELECT id_content, id_module, title, content, sequence, created_at, updated_at FROM practicum_module_content WHERE id_content IN (%s)", inClause)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	practicumModuleContents := []model.PracticumModuleContent{}
	for rows.Next() {
		practicumModuleContent := model.PracticumModuleContent{}
		if err := rows.Scan(&practicumModuleContent.IDContent, &practicumModuleContent.IDModule, &practicumModuleContent.Title, &practicumModuleContent.Content, &practicumModuleContent.Sequence, &practicumModuleContent.CreatedAt, &practicumModuleContent.UpdatedAt); err != nil {
			return nil, err
		}
		practicumModuleContents = append(practicumModuleContents, practicumModuleContent)
	}
	return practicumModuleContents, nil
}
