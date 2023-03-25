package postgresql

import (
	"app/api/models"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(req *models.CreateUser) (string, error) {
	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO users(
			id, 
			name, 
			balance,
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Balance,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *userRepo) GetByID(req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		user  models.User
	)

	query = `
		SELECT
			id,
			name,
			balance,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&user.Id,
		&user.Name,
		&user.Balance,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetList(req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

	resp = &models.GetListUserResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			balance,
			created_at,
			updated_at
		FROM users
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&resp.Count,
			&user.Id,
			&user.Name,
			&user.Balance,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return resp, nil
}

func (r *userRepo) Update(req *models.UpdateUser) (int64, error) {
	var (
		name    string
		balance string

		filter = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		users
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if req.Balance > 0 {
		balance = " balance = '" + strconv.Itoa(int(req.Balance)) + "', "
	}

	query += name + balance + " updated_at = now() " + filter

	result, err := r.db.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *userRepo) Delete(req *models.UserPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM users
		WHERE id = $1
	`

	result, err := r.db.Exec(query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
