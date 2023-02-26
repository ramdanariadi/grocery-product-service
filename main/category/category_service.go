package category

import "github.com/ramdanariadi/grocery-product-service/main/category/dto"

type CategoryService interface {
	FindAll(pageIndex int, pageSize int) *dto.AllCategories
	FindById(id string) *dto.CategoryDTO
	Save(body *dto.AddCategoryDTO)
	Update(id string, body *dto.AddCategoryDTO)
	Delete(id string)
}
