package usecase

import (
	"library-api/model"
	"library-api/repository"
)

type BookUseCase interface {
	GetAll() ([]model.Book, error)
	GetById(id int) (model.Book, error)
	Filter(authorID *int, categoryID *int) ([]model.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

// GetAll implements BookUseCase.
func (b *bookUseCase) GetAll() ([]model.Book, error) {
	return b.bookRepository.GetAll()
}

// GetById implements BookUseCase.
func (b *bookUseCase) GetById(id int) (model.Book, error) {
	return b.bookRepository.GetById(id)
}

// Filter implements BookUseCase.
func (b *bookUseCase) Filter(authorID *int, categoryID *int) ([]model.Book, error) {
	return b.bookRepository.Filter(authorID, categoryID)
}

func NewBookUsecase(repo repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: repo,
	}
}
