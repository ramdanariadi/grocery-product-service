package model

import (
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                 string               `json:"id" gorm:"primaryKey"`
	UserId             string               `json:"user_id"`
	TotalPrice         uint64               `json:"total_price"`
	TransactionDetails []*TransactionDetail `json:"detail_transaction"`
	CreatedAt          time.Time            `json:"transaction_date"`
	UpdatedAt          time.Time            `json:"_"`
	DeletedAt          gorm.DeletedAt       `json:"_" gorm:"index"`
}

type TransactionDetail struct {
	ID            string `json:"id" gorm:"primaryKey"`
	ProductId     string `json:"productId"`
	Product       product.Product
	TransactionId string `json:"transactionId"`
	Transaction   Transaction
	Price         uint64         `json:"price"`
	Weight        uint           `json:"weight"`
	Category      string         `json:"repository"`
	PerUnit       uint           `json:"perUnit"`
	Description   string         `json:"description"`
	ImageUrl      string         `json:"imageUrl"`
	Name          string         `json:"name"`
	CategoryId    string         `json:"categoryId"`
	Total         uint           `json:"total"`
	CreatedAt     time.Time      `json:"_"`
	UpdatedAt     time.Time      `json:"_"`
	DeletedAt     gorm.DeletedAt `json:"_" gorm:"index"`
}
