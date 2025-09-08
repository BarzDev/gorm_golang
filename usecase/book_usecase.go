package usecase

import (
	"fmt"
	"strconv"

	"library-api/model"
	"library-api/repository"
)

type BookUseCase interface {
	GetAll() ([]model.Book, error)
	GetById(id string) (model.Book, error)
	Filter(authorID *int, categoryID *int) ([]model.Book, error)
	Create(payload model.BookRequest) (model.Book, error)
	Update(id string, payload model.BookRequest) (model.Book, error)
	Delete(id string) error
	Validate(AuthorID int, CategoryID int) error
}

type bookUseCase struct {
	authorRepository   repository.AuthorRepository
	bookRepository     repository.BookRepository
	categoryRepository repository.CategoryRepository
}

// Validate implements BookUseCase.
func (b *bookUseCase) Validate(AuthorID int, CategoryID int) error {
	authorID := strconv.Itoa(AuthorID)
	categoryID := strconv.Itoa(CategoryID)

	_, err := b.authorRepository.GetById(authorID)
	if err != nil {
		return fmt.Errorf("author with id " + authorID + " not exist")
	}

	_, err = b.categoryRepository.GetById(categoryID)
	if err != nil {
		return fmt.Errorf("category with id " + categoryID + " not exist")
	}

	return nil
}

// Create implements BookUseCase.
func (b *bookUseCase) Create(payload model.BookRequest) (model.Book, error) {
	if err := b.Validate(payload.AuthorID, payload.CategoryID); err != nil {
		return model.Book{}, err
	}

	return b.bookRepository.Create(payload)
}

// Delete implements BookUseCase.
func (b *bookUseCase) Delete(id string) error {
	if err := b.bookRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Update implements BookUseCase.
func (b *bookUseCase) Update(id string, payload model.BookRequest) (model.Book, error) {
	if err := b.Validate(payload.AuthorID, payload.CategoryID); err != nil {
		return model.Book{}, err
	}

	return b.bookRepository.Update(id, payload)
}

// GetAll implements BookUseCase.
func (b *bookUseCase) GetAll() ([]model.Book, error) {
	return b.bookRepository.GetAll()
}

// GetById implements BookUseCase.
func (b *bookUseCase) GetById(id string) (model.Book, error) {
	return b.bookRepository.GetById(id)
}

// Filter implements BookUseCase.
func (b *bookUseCase) Filter(authorID *int, categoryID *int) ([]model.Book, error) {
	return b.bookRepository.Filter(authorID, categoryID)
}

func NewBookUsecase(
	authorRepository repository.AuthorRepository,
	bookRepository repository.BookRepository,
	categoryRepository repository.CategoryRepository,
) BookUseCase {
	return &bookUseCase{
		authorRepository:   authorRepository,
		bookRepository:     bookRepository,
		categoryRepository: categoryRepository,
	}
}
