package category

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        string         `gorm:"type:varchar(36);primaryKey"`
	Category  string         `gorm:"type:varchar(100);not null"`
	ImageUrl  string         `gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"_" gorm:"type:timestamp"`
	UpdatedAt time.Time      `json:"_" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"type:timestamp;index"`
}
