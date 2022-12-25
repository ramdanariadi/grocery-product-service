package model

type WishlistModel struct {
	Id        string
	Name      string
	Price     uint64
	Weight    uint32
	Category  string
	PerUnit   uint64
	ImageUrl  string
	ProductId string
	UserId    string
}
