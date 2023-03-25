package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"app/api/models"
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
			count,
			income_price,
			profit_status,
			profit_price,
			sell_price,
			author_id,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Count,
		req.IncomePrice,
		req.ProfitStatus,
		req.ProfitPrice,
		req.SellPrice,
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
			book.count,
			book.income_price,
			book.profit_status,
			book.profit_price,
			book.sell_price,
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
		&getBookRes.Count,
		&getBookRes.IncomePrice,
		&getBookRes.ProfitStatus,
		&getBookRes.ProfitPrice,
		&getBookRes.SellPrice,
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
			book.count,
			book.income_price,
			book.profit_status,
			book.profit_price,
			book.sell_price,
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
			&book.Count,
			&book.IncomePrice,
			&book.ProfitStatus,
			&book.ProfitPrice,
			&book.SellPrice,
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
		name         string
		incomePrice  string
		profitStatus string
		profitPrice  string
		sellPrice    string
		author_id    string
		filter       = " WHERE id = '" + req.Id + "'"
	)

	query := `
		UPDATE
		book
		SET
	`
	if len(req.Name) > 0 {
		name = " name = '" + req.Name + "', "
	}
	if req.Count > 0 {
		name = fmt.Sprintf(" count = %d ,", req.Count)
	}
	if req.IncomePrice > 0 {
		incomePrice = fmt.Sprintf(" income_price = %f ,", req.IncomePrice)
	}
	if len(req.ProfitStatus) > 0 {
		profitStatus = fmt.Sprintf(" profit_status = %s ,", req.ProfitStatus)
	}
	if req.ProfitPrice > 0 {
		profitPrice = fmt.Sprintf(" profit_price = %f ,", req.ProfitPrice)
	}
	if req.SellPrice > 0 {
		sellPrice = fmt.Sprintf(" sell_price = %f ,", req.SellPrice)
	}
	if len(req.AuthorId) > 0 {
		author_id = " author_id = '" + req.AuthorId + "', "
	}

	query += name + incomePrice + profitStatus + profitPrice + author_id + sellPrice + " updated_at = now() " + filter
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
