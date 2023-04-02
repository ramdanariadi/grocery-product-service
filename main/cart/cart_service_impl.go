package cart

import (
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/cart/dto"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"strings"
)

type ServiceImpl struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) Service {
	return &ServiceImpl{DB: DB}
}

func (service ServiceImpl) Store(productId string, total uint, userId string) {
	var productRef product.Product
	tx := service.DB.Where("id = ?", productId).Find(&productRef)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "INVALID_PRODUCT"})
	}

	id, _ := uuid.NewUUID()
	saveCart := Cart{
		ID:        id.String(),
		UserId:    userId,
		ProductId: productRef.ID,
		Total:     total,
	}
	save := service.DB.Create(&saveCart)
	utils.PanicIfError(save.Error)
}

func (service ServiceImpl) Destroy(id string, userId string) {
	cartRef := Cart{ID: id}
	tx := service.DB.Find(&cartRef)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "INVALID_CART"})
	}
	db := service.DB.Delete(&cartRef)
	utils.PanicIfError(db.Error)
}

func (service ServiceImpl) Find(reqBody *dto.FindCartDTO) []*dto.Cart {
	var carts []*Cart
	tx := service.DB.Model(&carts)
	tx.Joins("LEFT JOIN products p ON p.id = carts.product_id AND p.deleted_at IS NULL")
	tx.Joins("LEFT JOIN categories c ON p.category_id = c.id")
	tx.Preload("Product.Category")
	if reqBody.Search != nil {
		tx.Where("LOWER(p.name) LIKE ?", strings.ToLower("%"+*reqBody.Search+"%"))
	}
	tx.Limit(reqBody.PageSize).Offset(reqBody.PageIndex * reqBody.PageSize).Find(&carts)
	utils.PanicIfError(tx.Error)

	result := make([]*dto.Cart, 0)
	for _, data := range carts {
		result = append(result, &dto.Cart{
			ID:          data.ID,
			Total:       data.Total,
			Name:        data.Product.Name,
			Description: data.Product.Description,
			ImageUrl:    data.Product.ImageUrl,
			Price:       data.Product.Price,
			PerUnit:     data.Product.PerUnit,
			Weight:      data.Product.Weight,
			Category:    data.Product.Category.Category,
		})
	}

	return result
}
