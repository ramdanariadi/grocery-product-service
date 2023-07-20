package dto

type ProductDTO struct {
	ID          string `json:"id"`
	ShopId      string `json:"shopId"`
	ShopName    string `json:"shopName"`
	Price       uint64 `json:"price"`
	Weight      uint   `json:"weight"`
	Category    string `json:"category"`
	PerUnit     uint   `json:"perUnit"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Name        string `json:"name"`
}
