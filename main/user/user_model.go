package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id                string         `gorm:"type:varchar(36);primaryKey"`
	Username          string         `gorm:"type:varchar(100);uniqueIndex"`
	Password          string         `gorm:"type:varchar(255);not null"`
	Name              string         `gorm:"type:varchar(100);"`
	Email             string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	MobilePhoneNumber string         `gorm:"type:varchar(15);uniqueIndex"`
	CreatedAt         time.Time      `json:"_" gorm:"type:timestamp"`
	UpdatedAt         time.Time      `json:"_" gorm:"type:timestamp"`
	DeletedAt         gorm.DeletedAt `json:"_" gorm:"type:timestamp;index"`
}
