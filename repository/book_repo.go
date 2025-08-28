package repository

import (
	"library-api/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]model.Book, error)
	GetById(id int) (model.Book, error)
	Filter(authorID *int, categoryID *int) ([]model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

// GetAll implements BookRepository.
func (b *bookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book

	// Preload Author dan Category
	if err := b.db.Preload("Author").Preload("Category").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// GetById implements BookRepository.
func (b *bookRepository) GetById(id int) (model.Book, error) {
	var book model.Book

	// Preload Author dan Category
	err := b.db.Preload("Author").Preload("Category").First(&book, id).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// Filter implements BookRepository.
func (b *bookRepository) Filter(authorID *int, categoryID *int) ([]model.Book, error) {
	var books []model.Book
	db := b.db.Preload("Author").Preload("Category")

	if authorID != nil {
		db = db.Where("author_id = ?", *authorID)
	}
	if categoryID != nil {
		db = db.Where("category_id = ?", *categoryID)
	}

	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}
