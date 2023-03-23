package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"app/config"
	"app/storage"
)

type Store struct {
	db     *sql.DB
	book   storage.BookRepoI
	author storage.AuthorRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {

	connection := fmt.Sprintf(
		"host=%s user=%s database=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		db:     db,
		book:   NewBookRepo(db),
		author: NewAuthorRepo(db),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Book() storage.BookRepoI {

	if s.book == nil {
		s.book = NewBookRepo(s.db)
	}

	return s.book
}

func (s *Store) Author() storage.AuthorRepoI {
	if s.author == nil {
		s.author = NewAuthorRepo(s.db)
	}

	return s.author
}
