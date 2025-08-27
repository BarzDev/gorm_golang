package usecase

import (
	"library-api/model"
	"library-api/repository"
)

type AuthorUseCase interface {
	GetAll() ([]model.Author, error)
	GetById(id int) (model.Author, error)
}

type authorUsecase struct {
	authorRepository repository.AuthorRepository
}

// GetAll implements AuthorUseCase.
func (a *authorUsecase) GetAll() ([]model.Author, error) {
	return a.authorRepository.GetAll()
}

// GetById implements AuthorUseCase.
func (a *authorUsecase) GetById(id int) (model.Author, error) {
	return a.authorRepository.GetById(id)
}

func NewAuthorUsecase(repo repository.AuthorRepository) AuthorUseCase {
	return &authorUsecase{
		authorRepository: repo,
	}
}
