package service

import (
	"github.com/dan-ibm/go-crud-app/internal/domain"
)

type BooksRepository interface {
	Create(userId int, book domain.BookInput) error
	GetAll() ([]domain.Book, error)
	GetAllByUser(userId int) ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Delete(id int) error
	Update(id int, input domain.BookInput) error
}

type BookService struct {
	repo BooksRepository
}

func NewBookService(repo BooksRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (b *BookService) Create(userId int, book domain.BookInput) error {
	//if book.PublishDate.IsZero() {
	//	book.PublishDate = time.Now()
	//}

	return b.repo.Create(userId, book)
}

func (b *BookService) GetById(id int) (domain.Book, error) {
	return b.repo.GetById(id)
}

func (b *BookService) GetAll() ([]domain.Book, error) {
	return b.repo.GetAll()
}

func (b *BookService) GetAllByUser(userId int) ([]domain.Book, error) {
	return b.repo.GetAllByUser(userId)
}

func (b *BookService) Delete(id int) error {
	return b.repo.Delete(id)
}

func (b *BookService) Update(id int, inp domain.BookInput) error {
	return b.repo.Update(id, inp)
}
