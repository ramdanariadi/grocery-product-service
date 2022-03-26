package models

type CartModel struct {
	Id        string
	Name      string
	Price     int64
	Weight    uint
	Category  string
	PerUnit   int
	Total     int
	ImageUrl  interface{}
	ProductId string
	UserId    string
}
