package models

type WishlistModel struct {
	Id        string
	Name      string
	Price     int64
	Weight    uint
	Category  string
	PerUnit   int
	ImageUrl  string
	ProductId string
	UserId    string
}
