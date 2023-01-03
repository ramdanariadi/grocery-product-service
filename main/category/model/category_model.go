package model

import "time"

type CategoryModel struct {
	Id        string
	Category  string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
