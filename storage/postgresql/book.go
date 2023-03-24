package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"app/models"
)

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepo {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Create(req *models.CreateBook) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO book(
			id, 
			name, 
			price, 
			author_id,
			updated_at
		)
		VALUES ($1, $2, $3, $4, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Price,
		req.AuthorId,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *bookRepo) GetByID(req *models.BookPrimaryKey) (*models.GetBookRes, error) {

	var (
		query      string
		getBookRes models.GetBookRes
	)

	query = `
		SELECT
			book.id,
			book.name,
			book.price,
			book.created_at,
			book.updated_at,
			author.id,
			author.name,
			author.created_at,
			author.updated_at
		FROM book
		JOIN author ON author.id = book.author_id
		WHERE book.id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&getBookRes.Id,
		&getBookRes.Name,
		&getBookRes.Price,
		&getBookRes.CreatedAt,
		&getBookRes.UpdatedAt,
		&getBookRes.Author.Id,
		&getBookRes.Author.Name,
		&getBookRes.Author.UpdatedAt,
		&getBookRes.Author.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &getBookRes, nil
}

func (r *bookRepo) GetList(req *models.GetListBookRequest) (resp *models.GetListBookResponse, err error) {

	resp = &models.GetListBookResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			book.id,
			book.name,
			book.price,
			book.created_at,
			book.updated_at,
			author.id,
			author.name,
			author.created_at,
			author.updated_at
		FROM book
		JOIN author ON author.id = book.author_id
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

		var book models.GetBookRes
		err = rows.Scan(
			&resp.Count,
			&book.Id,
			&book.Name,
			&book.Price,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.Author.Id,
			&book.Author.Name,
			&book.Author.UpdatedAt,
			&book.Author.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Books = append(resp.Books, &book)
	}

	return resp, nil
}

func (r *bookRepo) Update(req *models.UpdateBook) (int64, error) {
	var (
		name      string
		price     string
		author_id string
		filter    = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		book
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if req.Price > 0 {
		price = fmt.Sprintf(" price = %f ,", req.Price)
	}
	if len(req.AuthorId) > 0 {
		author_id = " author_id = '" + req.AuthorId + "', "
	}

	query += name + price + author_id + " updated_at = now() " + filter
	fmt.Println(":::Query:", query)
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

func (r *bookRepo) Delete(req *models.BookPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM book
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

// Author
//   id
//   name

// Book
//   id
//   name
//   price
//   author_id = 1243256

// /book [GET]
//   {
//     "count": 1,
//     "books": [
//       {
//         "id": "ea4f35e9-c5ea-40b5-8b94-4db5c19c1243",
//         "name": "Learning golang",
//         "price": 400000,
//         "author": {
//           "id": "4e5657fe-d6ff-4fb3-8102-15d8c5b4efb3",
//           "name": "John Bodner"
//         }
//       }
//     ]
//   }
