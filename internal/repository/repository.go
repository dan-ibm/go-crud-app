package repository

import (
	"github.com/dan-ibm/go-crud-app/internal/domain"
	"github.com/dan-ibm/go-crud-app/internal/repository/psql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Book interface {
	Create(userId int, book domain.BookInput) error
	GetAll() ([]domain.Book, error)
	GetAllByUser(userId int) ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Delete(id int) error
	Update(id int, input domain.BookInput) error
}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psql.NewAuthRepository(db),
		Book:          psql.NewBookRepository(db),
	}
}
