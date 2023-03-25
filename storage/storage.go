package storage

import (
	"app/api/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	Author() AuthorRepoI
	User() UserRepoI
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

type UserRepoI interface {
	Create(req *models.CreateUser) (string, error)
	GetByID(req *models.UserPrimaryKey) (*models.User, error)
	GetList(req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
	Update(req *models.UpdateUser) (int64, error)
	Delete(req *models.UserPrimaryKey) (int64, error)
}
