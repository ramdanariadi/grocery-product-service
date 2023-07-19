package shop

type AddShopDTO struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
}

type EditShopDTO struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
}

type ShopDTO struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
