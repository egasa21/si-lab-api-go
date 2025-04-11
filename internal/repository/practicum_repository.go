package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type PracticumRepository interface {
	CreatePracticum(practicum *model.Practicum) error
	GetPracticumByID(id int) (*model.Practicum, error)
	GetPracticumByIDs(ids []int) ([]model.Practicum, error)
	GetAllPracticums(page, limit int) ([]model.Practicum, int, error)
	GetPracticumWithMaterialContents(id int) (*model.PracticumWithMaterial, error)
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

func (r *practicumRepository) GetPracticumByID(id int) (*model.Practicum, error) {
	var practicum model.Practicum
	err := r.db.QueryRow("SELECT id_practicum, name, code, description, credits, semester, created_at, updated_at FROM practicums WHERE id_practicum = $1", id).
		Scan(&practicum.ID, &practicum.Name, &practicum.Code, &practicum.Description, &practicum.Credits, &practicum.Semester, &practicum.CreatedAt, &practicum.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &practicum, nil
}

func (r *practicumRepository) GetAllPracticums(page, limit int) ([]model.Practicum, int, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Query(
		"SELECT id_practicum, name, code, description, credits, semester, created_at, updated_at FROM practicums LIMIT $1 OFFSET $2",
		limit, offset,
	)
	if err != nil {
		log.Error().Err(err)
		return nil, 0, err
	}

	defer rows.Close()

	var practicums []model.Practicum
	for rows.Next() {
		var practicum model.Practicum
		if err := rows.Scan(&practicum.ID, &practicum.Name, &practicum.Code, &practicum.Description, &practicum.Credits, &practicum.Semester, &practicum.CreatedAt, &practicum.UpdatedAt); err != nil {
			return nil, 0, err
		}
		practicums = append(practicums, practicum)
	}

	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM practicums").Scan(&total)
	if err != nil {
		log.Error().Err(err)
		return nil, 0, err
	}

	return practicums, total, nil
}

func (r *practicumRepository) GetPracticumByIDs(ids []int) ([]model.Practicum, error) {
	if len(ids) == 0 {
		return []model.Practicum{}, nil
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	inClause := strings.Join(placeholders, ",")
	query := fmt.Sprintf("SELECT id_practicum, name, code, description, credits, semester, created_at, updated_at FROM practicums WHERE id_practicum IN (%s)", inClause)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	practicums := []model.Practicum{}
	for rows.Next() {
		practicum := model.Practicum{}
		if err := rows.Scan(&practicum.ID, &practicum.Name, &practicum.Code, &practicum.Description, &practicum.Credits, &practicum.Semester, &practicum.CreatedAt, &practicum.UpdatedAt); err != nil {
			return nil, err
		}
		practicums = append(practicums, practicum)
	}

	return practicums, nil
}

func (r *practicumRepository) GetPracticumWithMaterialContents(id int) (*model.PracticumWithMaterial, error) {
	query := `
		SELECT 
			p.id_practicum, p.name, p.code, p.description, p.credits, p.semester,
			pm.id, pm.title,
			pmc.id_content, pmc.title
		FROM practicums p
		LEFT JOIN practicum_modules pm ON pm.practicum_id = p.id_practicum
		LEFT JOIN practicum_module_content pmc ON pmc.id_module = pm.id
		WHERE p.id_practicum = $1
		ORDER BY pm.id, pmc.sequence
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var practicum *model.PracticumWithMaterial
	moduleMap := map[uint]*model.ModuleWithMaterials{}

	for rows.Next() {
		var (
			practicumID                                  uint
			practicumName, code, desc, credits, semester string
			moduleID, contentID                          sql.NullInt64
			moduleTitle, contentTitle                    sql.NullString
		)

		if err := rows.Scan(
			&practicumID, &practicumName, &code, &desc, &credits, &semester,
			&moduleID, &moduleTitle,
			&contentID, &contentTitle,
		); err != nil {
			return nil, err
		}

		if practicum == nil {
			practicum = &model.PracticumWithMaterial{
				ID:          practicumID,
				Name:        practicumName,
				Code:        code,
				Description: desc,
				Credits:     credits,
				Semester:    semester,
				Modules:     []*model.ModuleWithMaterials{},
			}
		}

		if moduleID.Valid {
			modID := uint(moduleID.Int64)
			if _, exists := moduleMap[modID]; !exists {
				moduleMap[modID] = &model.ModuleWithMaterials{
					ID:        modID,
					Title:     moduleTitle.String,
					Materials: []model.Material{},
				}
				practicum.Modules = append(practicum.Modules, moduleMap[modID])
			}

			if contentID.Valid {
				moduleMap[modID].Materials = append(moduleMap[modID].Materials, model.Material{
					ID:    uint(contentID.Int64),
					Title: contentTitle.String,
				})
			}
		}
	}

	return practicum, nil
}
