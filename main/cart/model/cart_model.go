package model

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Price     uint64
	Weight    uint32
	Category  string
	PerUnit   uint64
	Total     uint32
	ImageUrl  string
	ProductId string
	UserId    string
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
