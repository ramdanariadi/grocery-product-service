package model

import (
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"gorm.io/gorm"
	"time"
)

type Wishlist struct {
	ID        string `gorm:"primaryKey"`
	ProductId string `gorm:"index"`
	Product   product.Product
	UserId    string         `gorm:"index"`
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
