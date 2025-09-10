package repository

import (
	"math"

	"library-api/model"
	"library-api/shared/shared_model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll(page, size int) ([]model.Book, shared_model.Paging, error)
	GetById(id string) (model.Book, error)
	Filter(authorID *int, categoryID *int) ([]model.Book, error)
	Create(payload model.BookRequest) (model.Book, error)
	Update(id string, payload model.BookRequest) (model.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

// Create implements BookRepository.
func (b *bookRepository) Create(payload model.BookRequest) (model.Book, error) {
	book := model.Book{
		Title:         payload.Title,
		AuthorID:      &payload.AuthorID,
		CategoryID:    &payload.CategoryID,
		PublishedYear: payload.PublishedYear,
		Price:         payload.Price,
		Stock:         payload.Stock,
	}

	if err := b.db.Create(&book).Error; err != nil {
		return model.Book{}, err
	}

	// Preload Author dan Category
	if err := b.db.Preload("Author").Preload("Category").First(&book, book.Id).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// Delete implements BookRepository.
func (b *bookRepository) Delete(id string) error {
	if err := b.db.Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Update implements BookRepository.
func (b *bookRepository) Update(id string, payload model.BookRequest) (model.Book, error) {
	var book model.Book

	if err := b.db.Model(&book).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(model.Book{
			Title:         payload.Title,
			AuthorID:      &payload.AuthorID,
			CategoryID:    &payload.CategoryID,
			PublishedYear: payload.PublishedYear,
			Price:         payload.Price,
			Stock:         payload.Stock,
		}).Error; err != nil {
		return model.Book{}, err
	}

	// Preload Author dan Category
	err := b.db.Preload("Author").Preload("Category").First(&book, id).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// GetAll implements BookRepository.
func (b *bookRepository) GetAll(page, size int) ([]model.Book, shared_model.Paging, error) {
	var books []model.Book
	var paging shared_model.Paging
	var totalRows int64

	if err := b.db.Model(&model.Book{}).Count(&totalRows).Error; err != nil {
		return nil, paging, err
	}

	offset := (page - 1) * size

	// Preload Author dan Category
	if err := b.db.
		Preload("Author").
		Preload("Category").
		Limit(size).
		Offset(offset).
		Find(&books).Error; err != nil {
		return nil, paging, err
	}

	paging = shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   int(totalRows),
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return books, paging, nil
}

// GetById implements BookRepository.
func (b *bookRepository) GetById(id string) (model.Book, error) {
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
