package product

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/product/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"strings"
)

type ProductServiceImpl struct {
	DB *gorm.DB
}

func (service ProductServiceImpl) Save(requestBody *dto.AddProductDTO) {
	var category category.Category
	tx := service.DB.Find(&category, "id = ?", requestBody.CategoryId)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product := Product{}
	id, err := uuid.NewUUID()
	utils.LogIfError(err)
	product.ID = id.String()
	product.Name = requestBody.Name
	product.ImageUrl = requestBody.ImageUrl
	product.Category = category
	product.Weight = requestBody.Weight
	product.Price = requestBody.Price
	product.PerUnit = requestBody.PerUnit
	product.IsRecommended = requestBody.IsRecommended
	product.IsTop = requestBody.IsTop
	product.Description = requestBody.Description
	save := service.DB.Create(&product)
	utils.LogIfError(save.Error)
}

func (service ProductServiceImpl) FindAll(param *dto.FindProductRequest) *dto.FindProductResponse {
	var products []Product
	tx := service.DB.Model(&Product{})

	if param.Search != nil {
		tx.Where("LOWER(name) like ?", strings.ToLower("%"+*param.Search+"%"))
	}

	if param.IsTop != nil {
		tx.Where("is_top = ?", *param.IsTop)
	}

	if param.IsRecommendation != nil {
		tx.Where("is_recommended = ?", param.IsRecommendation)
	}

	tx.Limit(param.PageSize).Offset(param.PageIndex * param.PageSize).Preload("Category").Find(&products)

	var result dto.FindProductResponse
	result.Data = make([]*dto.ProductDTO, 0)
	for _, p := range products {
		result.Data = append(result.Data, &dto.ProductDTO{
			ID:          p.ID,
			Name:        p.Name,
			ImageUrl:    p.ImageUrl,
			Category:    p.Category.Category,
			Weight:      p.Weight,
			Price:       p.Price,
			PerUnit:     p.PerUnit,
			Description: p.Description,
		})
	}
	result.RecordsFiltered = len(result.Data)
	var count int64
	tx.Count(&count)
	result.RecordsTotal = count
	return &result
}

func (service ProductServiceImpl) FindById(id string) *dto.ProductDTO {
	var result dto.ProductDTO
	var product Product
	tx := service.DB.Model(&Product{}).Where("id = ?", id).Preload("Category").Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	result.ID = product.ID
	result.Name = product.Name
	result.ImageUrl = product.ImageUrl
	result.Weight = product.Weight
	result.Price = product.Price
	result.PerUnit = product.PerUnit
	result.Description = product.Description
	result.Category = product.Category.Category
	return &result
}

func (service ProductServiceImpl) Update(id string, requestBody *dto.AddProductDTO) {
	var category category.Category
	txCategory := service.DB.Find(&category, "id = ?", requestBody.CategoryId)
	if txCategory.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	var product Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product.Name = requestBody.Name
	product.ImageUrl = requestBody.ImageUrl
	product.Category = category
	product.Weight = requestBody.Weight
	product.Price = requestBody.Price
	product.PerUnit = requestBody.PerUnit
	product.IsRecommended = requestBody.IsRecommended
	product.IsTop = requestBody.IsTop
	product.Description = requestBody.Description
	service.DB.Save(&product)
}

func (service ProductServiceImpl) Delete(id string) {
	var product Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	service.DB.Where("id = ?", id).Delete(&product)
}

func (service ProductServiceImpl) SetTop(id string) {
	var product Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product.IsTop = true
	service.DB.Save(&product)
}

func (service ProductServiceImpl) SetRecommendation(id string) {
	var product Product
	tx := service.DB.Where("id = ?", id).Find(&product)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	product.IsRecommended = true
	service.DB.Save(&product)
}
