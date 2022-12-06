package models

type ProductModel struct {
	Id            string      `json:"id"`
	Price         uint64      `json:"price"`
	Weight        uint32      `json:"weight"`
	Category      string      `json:"category"`
	PerUnit       uint64      `json:"perUnit"`
	Description   string      `json:"description"`
	ImageUrl      interface{} `json:"imageUrl"`
	Name          string      `json:"name"`
	CategoryId    string      `json:"categoryId"`
	IsRecommended bool        `json:"isRecommended"`
	IsTop         bool        `json:"isTop"`
}
