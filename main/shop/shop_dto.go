package shop

type AddShopDTO struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	ImageUrl *string `json:"imageUrl"`
}

type EditShopDTO struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	ImageUrl *string `json:"imageUrl"`
}

type ShopDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	ImageUrl string `json:"imageUrl"`
}
