package dto

type TransactionDTO struct {
	Id         string                `json:"id"`
	Date       string                `json:"date"`
	PriceTotal uint64                `json:"priceTotal"`
	Items      []*TransactionItemDTO `json:"items"`
}

type TransactionItemDTO struct {
	ID          string `json:"id"`
	Price       uint64 `json:"price"`
	Weight      uint   `json:"weight"`
	PerUnit     uint   `json:"perUnit"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Name        string `json:"name"`
	Total       uint   `json:"total"`
}
