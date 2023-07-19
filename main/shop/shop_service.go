package shop

type ShopService interface {
	AddShop(userId string, shop AddShopDTO)
	UpdateShop(userId string, shop EditShopDTO)
	GetShop(userId string) ShopDTO
	DeleteShop(userID string)
}
