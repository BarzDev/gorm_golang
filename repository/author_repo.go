package repository

import (
	"library-api/model"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAll() ([]model.Author, error)
	GetById(id int) (model.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

// GetAll implements AuthorRepository.
func (a *authorRepository) GetAll() ([]model.Author, error) {
	var authors []model.Author
	if err := a.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

// GetById implements AuthorRepository.
func (a *authorRepository) GetById(id int) (model.Author, error) {
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
