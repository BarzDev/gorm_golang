package repository

import (
	"library-api/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetById(id int) (model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

// GetAll implements CategoryRepository.
func (c *categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetById implements CategoryRepository.
func (c *categoryRepository) GetById(id int) (model.Category, error) {
	var category model.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
