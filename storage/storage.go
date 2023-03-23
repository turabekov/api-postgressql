package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	Author() AuthorRepoI
}

type BookRepoI interface {
	Create(*models.CreateBook) (string, error)
	GetByID(*models.BookPrimaryKey) (*models.GetBookRes, error)
	GetList(*models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(*models.Book) (*models.BookPrimaryKey, error)
	Delete(req *models.BookPrimaryKey) (err error)
}

type AuthorRepoI interface {
	Create(req *models.CreateAuthor) (string, error)
	GetByID(req *models.AuthorPrimaryKey) (*models.Author, error)
	GetList(req *models.GetListAuthorRequest) (resp *models.GetListAuthorResponse, err error)
	Update(req *models.UpdateAuthor) (res *models.AuthorPrimaryKey, err error)
	Delete(req *models.AuthorPrimaryKey) (err error)
}
