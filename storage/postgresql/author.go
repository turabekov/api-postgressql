package postgresql

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type authorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *authorRepo {
	return &authorRepo{
		db: db,
	}
}
func (r *authorRepo) Create(req *models.CreateAuthor) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO author(
			id, 
			name, 
			updated_at
		)
		VALUES ($1, $2, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *authorRepo) GetByID(req *models.AuthorPrimaryKey) (*models.Author, error) {

	var (
		query  string
		author models.Author
	)

	query = `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM author
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&author.Id,
		&author.Name,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *authorRepo) GetList(req *models.GetListAuthorRequest) (resp *models.GetListAuthorResponse, err error) {

	resp = &models.GetListAuthorResponse{}

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
			created_at,
			updated_at
		FROM author
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

	fmt.Println(":::Query:", query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var author models.Author
		err = rows.Scan(
			&resp.Count,
			&author.Id,
			&author.Name,
			&author.CreatedAt,
			&author.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Authors = append(resp.Authors, &author)
	}

	return resp, nil
}

func (r *authorRepo) Update(req *models.UpdateAuthor) (res *models.AuthorPrimaryKey, err error) {
	var (
		name   string
		filter = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		author
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}

	query += name + " updated_at = now() " + filter
	fmt.Println(":::Query:", query)
	_, err = r.db.Exec(query)

	if err != nil {
		return nil, err
	}

	id := models.AuthorPrimaryKey{
		Id: req.Id,
	}

	return &id, nil
}

func (r *authorRepo) Delete(req *models.AuthorPrimaryKey) (err error) {
	query := `
		DELETE 
		FROM author
		WHERE id = $1
	`

	_, err = r.db.Exec(query, req.Id)

	fmt.Println(":::Query:", query)
	if err != nil {
		return err
	}

	return nil
}
