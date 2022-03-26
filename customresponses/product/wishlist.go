package product

type WishlistResponse struct {
	Id       string      `json:"id"`
	Price    int64       `json:"price"`
	Weight   uint        `json:"weight"`
	Category string      `json:"category"`
	PerUnit  int         `json:"perUnit"`
	ImageUrl interface{} `json:"imageUrl"`
	Name     string      `json:"name"`
}
