package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        string `gorm:"primaryKey"`
	Category  string
	ImageUrl  string
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
