package model

type CartModel struct {
	Id        string
	Name      string
	Price     uint64
	Weight    uint32
	Category  string
	PerUnit   uint64
	Total     uint32
	ImageUrl  interface{}
	ProductId string
	UserId    string
}
