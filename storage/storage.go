package storage

import (
	"app/api/models"
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
	Update(req *models.UpdateBook) (int64, error)
	Delete(req *models.BookPrimaryKey) (int64, error)
}

type AuthorRepoI interface {
	Create(req *models.CreateAuthor) (string, error)
	GetByID(req *models.AuthorPrimaryKey) (*models.Author, error)
	GetList(req *models.GetListAuthorRequest) (resp *models.GetListAuthorResponse, err error)
	Update(req *models.UpdateAuthor) (int64, error)
	Delete(req *models.AuthorPrimaryKey) (int64, error)
}
