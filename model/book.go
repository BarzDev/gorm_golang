package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id            int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string         `json:"title"`
	AuthorID      *int           `json:"-"` // foreign key
	Author        Author         `json:"author"`
	CategoryID    *int           `json:"-"` // foreign key
	Category      Category       `json:"category"`
	PublishedYear int            `json:"published_year"`
	Price         float64        `json:"price"`
	Stock         int            `json:"stock"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type BookRequest struct {
	Title         string  `json:"title" binding:"required"`
	AuthorID      int     `json:"author_id" binding:"required"`
	CategoryID    int     `json:"category_id" binding:"required"`
	PublishedYear int     `json:"published_year" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
}
