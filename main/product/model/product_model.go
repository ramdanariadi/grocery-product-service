package model

import (
	"github.com/ramdanariadi/grocery-product-service/main/category/model"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID            string `json:"id" gorm:"primaryKey"`
	Price         uint64 `json:"price"`
	Weight        uint   `json:"weight"`
	CategoryId    string
	Category      model.Category
	PerUnit       uint           `json:"perUnit"`
	Description   string         `json:"description"`
	ImageUrl      string         `json:"imageUrl"`
	Name          string         `json:"name"`
	IsRecommended bool           `json:"isRecommended"`
	IsTop         bool           `json:"isTop"`
	CreatedAt     time.Time      `json:"_"`
	UpdatedAt     time.Time      `json:"_"`
	DeletedAt     gorm.DeletedAt `json:"_" gorm:"index"`
}
