package dto

type AddCategoryDTO struct {
	Category string `json:"category"`
	ImageUrl string `bson:"imageUrl"`
}
