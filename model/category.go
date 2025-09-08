package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
