package usecase

import (
	"library-api/model"
	"library-api/repository"
)

type AuthorUseCase interface {
	GetAll() ([]model.Author, error)
	GetById(id string) (model.Author, error)
	Create(payload model.AuthorRequest) (model.Author, error)
	Update(id string, payload model.AuthorRequest) (model.Author, error)
	Delete(id string) error
}

type authorUsecase struct {
	authorRepository repository.AuthorRepository
}

// Create implements AuthorUseCase.
func (a *authorUsecase) Create(payload model.AuthorRequest) (model.Author, error) {
	author, err := a.authorRepository.Create(payload)
	if err != nil {
		return model.Author{}, err
	}
	return author, nil
}

// Delete implements AuthorUseCase.
func (a *authorUsecase) Delete(id string) error {
	if err := a.authorRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Update implements AuthorUseCase.
func (a *authorUsecase) Update(id string, payload model.AuthorRequest) (model.Author, error) {
	updatedAuthor, err := a.authorRepository.Update(id, payload)
	if err != nil {
		return model.Author{}, err
	}
	return updatedAuthor, nil
}

// GetAll implements AuthorUseCase.
func (a *authorUsecase) GetAll() ([]model.Author, error) {
	return a.authorRepository.GetAll()
}

// GetById implements AuthorUseCase.
func (a *authorUsecase) GetById(id string) (model.Author, error) {
	return a.authorRepository.GetById(id)
}

func NewAuthorUsecase(repo repository.AuthorRepository) AuthorUseCase {
	return &authorUsecase{
		authorRepository: repo,
	}
}
