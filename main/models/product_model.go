package models

type ProductModel struct {
	Id            string `json:"id"`
	Price         uint64 `json:"price"`
	Weight        uint   `json:"weight"`
	Category      string `json:"category"`
	PerUnit       uint   `json:"perUnit"`
	Description   string `json:"description"`
	ImageUrl      string `json:"imageUrl"`
	Name          string `json:"name"`
	CategoryId    string `json:"categoryId"`
	IsRecommended bool   `json:"isRecommended"`
	IsTop         bool   `json:"isTop"`
}
