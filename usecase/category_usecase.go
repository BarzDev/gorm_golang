package usecase

import (
	"library-api/model"
	"library-api/repository"
)

type CategoryUseCase interface {
	GetAll() ([]model.Category, error)
	GetById(id string) (model.Category, error)
	Create(payload model.CategoryRequest) (model.Category, error)
	Update(id string, payload model.CategoryRequest) (model.Category, error)
	Delete(id string) error
}

type categoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

// Create implements CategoryUseCase.
func (c *categoryUseCase) Create(payload model.CategoryRequest) (model.Category, error) {
	category, err := c.categoryRepository.Create(payload)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// Delete implements CategoryUseCase.
func (c *categoryUseCase) Delete(id string) error {
	if err := c.categoryRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Update implements CategoryUseCase.
func (c *categoryUseCase) Update(id string, payload model.CategoryRequest) (model.Category, error) {
	updateCategory, err := c.categoryRepository.Update(id, payload)
	if err != nil {
		return model.Category{}, err
	}
	return updateCategory, nil
}

// GetAll implements CategoryUseCase.
func (c *categoryUseCase) GetAll() ([]model.Category, error) {
	return c.categoryRepository.GetAll()
}

// GetById implements CategoryUseCase.
func (c *categoryUseCase) GetById(id string) (model.Category, error) {
	return c.categoryRepository.GetById(id)
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: repo,
	}
}
