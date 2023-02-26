package category

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        string         `gorm:"type:varchar(36);primaryKey"`
	Category  string         `gorm:"type:varchar(100)"`
	ImageUrl  string         `gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
