package product

import "github.com/ramdanariadi/grocery-product-service/main/product/dto"

type ProductService interface {
	Save(product *dto.AddProductDTO)
	FindAll(param *dto.FindProductRequest) *dto.FindProductResponse
	FindById(id string) *dto.ProductDTO
	Update(id string, product *dto.AddProductDTO)
	Delete(id string)
	SetTop(id string)
	SetRecommendation(id string)
}
