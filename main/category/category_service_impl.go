package category

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/category/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"log"
)

type CategoryServiceImpl struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func (service CategoryServiceImpl) FindAll(pageIndex int, pageSize int) *dto.AllCategories {
	var categories []*Category
	service.DB.Limit(pageSize).Offset(pageSize * pageIndex).Find(&categories)
	var count int64
	service.DB.Model(&Category{}).Where("deleted_at IS NULL").Count(&count)

	result := dto.AllCategories{}
	result.Data = make([]*dto.CategoryDTO, 0)

	for _, category := range categories {
		result.Data = append(result.Data, &dto.CategoryDTO{Id: category.ID, Category: category.Category, ImageUrl: category.ImageUrl})
	}

	result.RecordsFiltered = len(result.Data)
	result.RecordsTotal = count
	return &result
}

func (service CategoryServiceImpl) FindById(id string) *dto.CategoryDTO {
	var category Category
	var result dto.CategoryDTO
	tx := service.DB.Where("id = ? AND deleted_at IS NULL", id).Find(&category)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	result.Id = category.ID
	result.Category = category.Category
	result.ImageUrl = category.ImageUrl
	return &result
}

func (service CategoryServiceImpl) Save(body *dto.AddCategoryDTO) {
	id, err := uuid.NewUUID()
	utils.LogIfError(err)
	service.DB.Create(&Category{ID: id.String(), Category: body.Category, ImageUrl: body.ImageUrl})
}

func (service CategoryServiceImpl) Update(id string, body *dto.AddCategoryDTO) {
	var category Category
	tx := service.DB.Where("id = ?", id).Find(&category)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}
	category.Category = body.Category
	category.ImageUrl = body.ImageUrl
	log.Printf("Id %s", id)
	save := service.DB.Save(&category)
	if save.Error != nil {
		panic(exception.InternalServerError)
	}
}

func (service CategoryServiceImpl) Delete(id string) {
	var category Category
	tx := service.DB.Where("id = ?", id).Find(&category)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: exception.BadRequest})
	}

	service.DB.Delete(&category)
}
