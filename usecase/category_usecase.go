package usecase

import (
	"library-api/model"
	"library-api/repository"
)

type CategoryUseCase interface {
	GetAll() ([]model.Category, error)
	GetById(id int) (model.Category, error)
}

type categoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

// GetAll implements CategoryUseCase.
func (c *categoryUseCase) GetAll() ([]model.Category, error) {
	return c.categoryRepository.GetAll()
}

// GetById implements CategoryUseCase.
func (c *categoryUseCase) GetById(id int) (model.Category, error) {
	return c.categoryRepository.GetById(id)
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: repo,
	}
}
