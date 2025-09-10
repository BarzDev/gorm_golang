package repository

import (
	"math"

	"library-api/model"
	"library-api/shared/shared_model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(page, size int) ([]model.Category, shared_model.Paging, error)
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
func (c *categoryRepository) GetAll(page, size int) ([]model.Category, shared_model.Paging, error) {
	var categories []model.Category
	var paging shared_model.Paging
	var totalRows int64

	if err := c.db.Model(&model.Author{}).Count(&totalRows).Error; err != nil {
		return nil, paging, err
	}

	offset := (page - 1) * size

	if err := c.db.
		Limit(size).
		Offset(offset).
		Find(&categories).Error; err != nil {
		return nil, paging, err
	}

	paging = shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   int(totalRows),
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return categories, paging, nil
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
