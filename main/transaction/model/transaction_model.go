package model

import (
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                 string `gorm:"primaryKey"`
	UserId             string
	TotalPrice         uint64
	TransactionDetails []*TransactionDetail
	CreatedAt          time.Time      `json:"_"`
	UpdatedAt          time.Time      `json:"_"`
	DeletedAt          gorm.DeletedAt `json:"_" gorm:"index"`
}

type TransactionDetail struct {
	ID            string `gorm:"primaryKey"`
	ProductId     string
	Product       product.Product
	TransactionId string
	Transaction   Transaction
	Price         uint64
	Weight        uint
	Category      string
	PerUnit       uint
	Description   string
	ImageUrl      string
	Name          string
	CategoryId    string
	Total         uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
