package models

type WishlistModel struct {
	Id        string
	Name      string
	Price     uint64
	Weight    uint32
	Category  string
	PerUnit   uint64
	ImageUrl  interface{}
	ProductId string
	UserId    string
}
