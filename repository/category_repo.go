package repository

import (
	"library-api/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	GetById(id string) (model.Category, error)
	Create(payload model.CategoryRequest) (model.Category, error)
	Update(id string, payload model.CategoryRequest) (model.Category, error)
	Delete(id string) error
}

type categoryRepository struct {
	db *gorm.DB
}

// Create implements CategoryRepository.
func (c *categoryRepository) Create(payload model.CategoryRequest) (model.Category, error) {
	category := model.Category{
		Name: payload.Name,
	}

	if err := c.db.Create(&category).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// Delete implements CategoryRepository.
func (c *categoryRepository) Delete(id string) error {
	if err := c.db.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Update implements CategoryRepository.
func (c *categoryRepository) Update(id string, payload model.CategoryRequest) (model.Category, error) {
	var category model.Category

	if err := c.db.Model(&category).Where("id = ? AND deleted_at IS NULL", id).Updates(map[string]interface{}{
		"name": payload.Name,
	}).Error; err != nil {
		return model.Category{}, err
	}

	if err := c.db.First(&category, id).Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
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
func (c *categoryRepository) GetById(id string) (model.Category, error) {
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
