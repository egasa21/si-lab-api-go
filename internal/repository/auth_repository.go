package repository

import (
	"database/sql"
	"fmt"

	"github.com/egasa21/si-lab-api-go/internal/model"
)

type AuthRepository interface {
	Register(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	AddRoleToUser(idUser int, idRole int) error
	GetRolesByUserID(idUser int) ([]model.RoleModel, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(user *model.User) error {
	query := `
		INSERT INTO users (email, password, id_student)
		VALUES ($1, $2, $3)
		RETURNING id_user`

	// Scan the returned id_user into the user struct
	err := r.db.QueryRow(query, user.Email, user.Password, user.IDStudent).Scan(&user.IDUser)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}

func (r *authRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `
		SELECT id_user, email, password, id_student, created_at
		FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.IDUser,
		&user.Email,
		&user.Password,
		&user.IDStudent,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	roles, err := r.GetRolesByUserID(user.IDUser)
	if err != nil {
		return nil, err
	}
	user.Roles = roles

	return &user, nil
}

func (r *authRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	query := `
		SELECT id_user, email, password, id_student, created_at
		FROM users WHERE id_user = $1`
	err := r.db.QueryRow(query, id).Scan(
		&user.IDUser,
		&user.Email,
		&user.Password,
		&user.IDStudent,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	roles, err := r.GetRolesByUserID(user.IDUser)
	if err != nil {
		return nil, err
	}
	user.Roles = roles

	return &user, nil
}

func (r *authRepository) AddRoleToUser(idUser int, idRole int) error {
	query := `
		INSERT INTO user_roles (id_user, id_role)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`
	_, err := r.db.Exec(query, idUser, idRole)
	return err
}

func (r *authRepository) GetRolesByUserID(idUser int) ([]model.RoleModel, error) {
	query := `
	SELECT r.id, r.name
	FROM roles r
	INNER JOIN user_roles ur ON r.id = ur.id_role
	WHERE ur.id_user = $1`

	rows, err := r.db.Query(query, idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.RoleModel
	for rows.Next() {
		var role model.RoleModel
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
