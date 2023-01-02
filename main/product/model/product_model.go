package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID            string         `json:"id" gorm:"primaryKey"`
	Price         uint64         `json:"price"`
	Weight        uint           `json:"weight"`
	Category      string         `json:"repository"`
	PerUnit       uint           `json:"perUnit"`
	Description   string         `json:"description"`
	ImageUrl      string         `json:"imageUrl"`
	Name          string         `json:"name"`
	CategoryId    string         `json:"categoryId"`
	IsRecommended bool           `json:"isRecommended"`
	IsTop         bool           `json:"isTop"`
	CreatedAt     time.Time      `json:"_"`
	UpdatedAt     time.Time      `json:"_"`
	DeletedAt     gorm.DeletedAt `json:"_" gorm:"index"`
}
