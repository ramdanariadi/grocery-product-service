package shop

import (
	"github.com/ramdanariadi/grocery-product-service/main/user"
	"gorm.io/gorm"
	"time"
)

type Shop struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	ImageUrl  string `gorm:"type:varchar(255);not null"`
	Address   string `gorm:"type:varchar(300)"`
	UserId    string `gorm:"index"`
	User      user.User
	CreatedAt time.Time      `json:"_"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
