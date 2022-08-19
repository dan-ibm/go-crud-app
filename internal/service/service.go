package service

import (
	"github.com/dan-ibm/go-crud-app/internal/domain"
	"github.com/dan-ibm/go-crud-app/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Book interface {
	Create(userId int, book domain.BookInput) error
	GetAll() ([]domain.Book, error)
	GetAllByUser(userId int) ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Delete(id int) error
	Update(id int, input domain.BookInput) error
}

type Service struct {
	Authorization
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Book:          NewBookService(repos.Book),
	}
}
