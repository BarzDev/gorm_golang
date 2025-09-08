package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	Id        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Bio       *string        `json:"bio"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AuthorRequest struct {
	Name string  `json:"name" binding:"required"`
	Bio  *string `json:"bio"`
}
