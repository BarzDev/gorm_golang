package repository

import (
	"math"

	"library-api/model"
	"library-api/shared/shared_model"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAll(page, size int) ([]model.Author, shared_model.Paging, error)
	GetById(id string) (model.Author, error)
	Create(payload model.AuthorRequest) (model.Author, error)
	Update(id string, payload model.AuthorRequest) (model.Author, error)
	Delete(id string) error
}

type authorRepository struct {
	db *gorm.DB
}

// Create implements AuthorRepository.
func (a *authorRepository) Create(payload model.AuthorRequest) (model.Author, error) {
	author := model.Author{
		Name: payload.Name,
		Bio:  payload.Bio,
	}

	if err := a.db.Create(&author).Error; err != nil {
		return model.Author{}, err
	}
	return author, nil
}

// Delete implements AuthorRepository.
func (a *authorRepository) Delete(id string) error {
	if err := a.db.Delete(&model.Author{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Update implements AuthorRepository.
func (a *authorRepository) Update(id string, payload model.AuthorRequest) (model.Author, error) {
	var author model.Author

	if err := a.db.Model(&author).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"name": payload.Name,
			"bio":  payload.Bio,
		}).Error; err != nil {
		return model.Author{}, err
	}

	if err := a.db.First(&author, id).Error; err != nil {
		return model.Author{}, err
	}

	return author, nil
}

// GetAll implements AuthorRepository.
func (a *authorRepository) GetAll(page, size int) ([]model.Author, shared_model.Paging, error) {
	var authors []model.Author
	var paging shared_model.Paging
	var totalRows int64

	if err := a.db.Model(&model.Author{}).Count(&totalRows).Error; err != nil {
		return nil, paging, err
	}

	offset := (page - 1) * size

	if err := a.db.
		Limit(size).
		Offset(offset).
		Find(&authors).Error; err != nil {
		return nil, paging, err
	}

	paging = shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   int(totalRows),
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return authors, paging, nil
}

// GetById implements AuthorRepository.
func (a *authorRepository) GetById(id string) (model.Author, error) {
	var author model.Author
	if err := a.db.First(&author, id).Error; err != nil {
		return model.Author{}, err
	}
	return author, nil
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{
		db: db,
	}
}
