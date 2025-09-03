package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
