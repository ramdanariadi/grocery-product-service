package dto

type AddProductDTO struct {
	Price         uint64 `json:"price"`
	Weight        uint   `json:"weight"`
	CategoryId    string `json:"categoryId"`
	PerUnit       uint   `json:"perUnit"`
	Description   string `json:"description"`
	ImageUrl      string `json:"imageUrl"`
	Name          string `json:"name"`
	IsRecommended bool   `json:"isRecommended"`
	IsTop         bool   `json:"isTop"`
}
