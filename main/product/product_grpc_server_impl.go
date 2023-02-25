package product

//
//import (
//	"encoding/json"
//	"github.com/go-redis/redis/v8"
//	"github.com/google/uuid"
//	"github.com/ramdanariadi/grocery-product-service/main/category"
//	categoryModel "github.com/ramdanariadi/grocery-product-service/main/category/model"
//	"github.com/ramdanariadi/grocery-product-service/main/product/model"
//	"github.com/ramdanariadi/grocery-product-service/main/response"
//	"github.com/ramdanariadi/grocery-product-service/main/setup"
//	"github.com/ramdanariadi/grocery-product-service/main/utils"
//	"golang.org/x/net/context"
//	"gorm.io/gorm"
//	"log"
//	"time"
//)
//
//type ProductServiceServerImpl struct {
//	DB          *gorm.DB
//	RedisClient *redis.Client
//}
//
//func NewProductServiceServerImpl(db *gorm.DB) *ProductServiceServerImpl {
//	return &ProductServiceServerImpl{
//		DB:          db,
//		RedisClient: setup.NewRedisClient(),
//	}
//}
//
//func (server ProductServiceServerImpl) FindById(ctx context.Context, id *ProductId) (*ProductResponse, error) {
//	var productModel *model.Product
//	cache, err := server.RedisClient.Get(ctx, id.GetId()).Result()
//	utils.LogIfError(err)
//
//	if cache != "" {
//		err := json.Unmarshal([]byte(cache), &productModel)
//		utils.LogIfError(err)
//	} else {
//		tx := server.DB.Find(&productModel, "id = ?", id.Id)
//		log.Printf("find rows affected : %d", tx.RowsAffected)
//		if tx.RowsAffected == 0 {
//			status, message := utils.QueryResponse(false)
//			return &ProductResponse{
//				Message: message,
//				Status:  status,
//				Data:    nil,
//			}, nil
//		}
//		bytes, err := json.Marshal(productModel)
//		utils.LogIfError(err)
//		err = server.RedisClient.Set(ctx, id.GetId(), bytes, 1*time.Hour).Err()
//		utils.LogIfError(err)
//	}
//
//	grpcProductModel := Product{
//		Id:          productModel.ID,
//		Name:        productModel.Name,
//		Weight:      uint32(productModel.Weight),
//		Category:    productModel.Category.Category,
//		ImageUrl:    productModel.ImageUrl,
//		CategoryId:  productModel.CategoryId,
//		Price:       productModel.Price,
//		PerUnit:     uint64(productModel.PerUnit),
//		Description: productModel.Description,
//	}
//	status, message := utils.QueryResponse(true)
//	return &ProductResponse{
//		Message: message,
//		Status:  status,
//		Data:    &grpcProductModel,
//	}, nil
//}
//
//func (server ProductServiceServerImpl) FindProductsByCategory(_ context.Context, id *category.CategoryId) (*MultipleProductResponse, error) {
//	var productsModel []*model.Product
//	tx := server.DB.Find(&productsModel, "category_id = ?", id.Id)
//	utils.PanicIfError(tx.Error)
//
//	products := fetchProducts(productsModel)
//	status, message := utils.QueryResponse(len(products) > 0)
//	return &MultipleProductResponse{
//		Status:  status,
//		Data:    products,
//		Message: message,
//	}, nil
//}
//
//func (server ProductServiceServerImpl) FindRecommendedProduct(_ context.Context, _ *ProductEmpty) (*MultipleProductResponse, error) {
//	var recommendationProducts []*model.Product
//	tx := server.DB.Find(&recommendationProducts, "products.is_recommended = ?", true)
//	utils.PanicIfError(tx.Error)
//	products := fetchProducts(recommendationProducts)
//	status, message := utils.QueryResponse(len(products) > 0)
//	return &MultipleProductResponse{
//		Status:  status,
//		Data:    products,
//		Message: message,
//	}, nil
//}
//
//func (server ProductServiceServerImpl) FindTopProducts(_ context.Context, _ *ProductEmpty) (*MultipleProductResponse, error) {
//	var recommendationProducts []*model.Product
//	tx := server.DB.Find(&recommendationProducts, "products.is_top = ?", true)
//	utils.PanicIfError(tx.Error)
//	products := fetchProducts(recommendationProducts)
//	status, message := utils.QueryResponse(len(products) > 0)
//	return &MultipleProductResponse{
//		Status:  status,
//		Data:    products,
//		Message: message,
//	}, nil
//}
//
//func (server ProductServiceServerImpl) FindAll(_ context.Context, _ *ProductEmpty) (*MultipleProductResponse, error) {
//	var recommendationProducts []*model.Product
//	tx := server.DB.Find(&recommendationProducts)
//	utils.PanicIfError(tx.Error)
//	products := fetchProducts(recommendationProducts)
//	status, message := utils.QueryResponse(true)
//	return &MultipleProductResponse{
//		Status:  status,
//		Data:    products,
//		Message: message,
//	}, nil
//}
//
//func fetchProducts(productsModel []*model.Product) []*Product {
//	var products []*Product
//	for _, p := range productsModel {
//		productTmp := Product{
//			Id:            p.ID,
//			Name:          p.Name,
//			Weight:        uint32(p.Weight),
//			Price:         p.Price,
//			Description:   p.Description,
//			ImageUrl:      p.ImageUrl,
//			PerUnit:       uint64(p.PerUnit),
//			Category:      p.Category.Category,
//			CategoryId:    p.CategoryId,
//			IsTop:         p.IsTop,
//			IsRecommended: p.IsRecommended,
//		}
//		products = append(products, &productTmp)
//	}
//	return products
//}
//
//func (server ProductServiceServerImpl) Save(_ context.Context, product *Product) (*response.Response, error) {
//	categoryReff := categoryModel.Category{ID: product.CategoryId}
//	tx := server.DB.First(&categoryReff)
//	if tx.RowsAffected == 0 {
//		status, _ := utils.QueryResponse(false)
//		return &response.Response{Status: status, Message: "INVALID_CATEGORY"}, nil
//	}
//
//	id, _ := uuid.NewUUID()
//	productModel := model.Product{
//		ID:          id.String(),
//		Name:        product.Name,
//		Weight:      uint(product.Weight),
//		Category:    categoryReff,
//		CategoryId:  categoryReff.ID,
//		Price:       product.Price,
//		PerUnit:     uint(product.PerUnit),
//		Description: product.Description,
//		ImageUrl:    product.ImageUrl,
//	}
//	save := server.DB.Save(&productModel)
//	status, message := utils.ModifyingResponse(save.Error == nil)
//	return &response.Response{
//		Status:  status,
//		Message: message,
//	}, nil
//}
//
//func (server ProductServiceServerImpl) Update(ctx context.Context, product *Product) (*response.Response, error) {
//	server.RedisClient.Del(ctx, product.Id)
//	var productModel model.Product
//	find := server.DB.Find(&productModel, "id = ?", product.Id)
//	if find.RowsAffected == 0 {
//		status, _ := utils.QueryResponse(false)
//		return &response.Response{Status: status, Message: "INVALID_PRODUCT"}, nil
//	}
//
//	categoryReff := categoryModel.Category{ID: product.CategoryId}
//	tx := server.DB.First(&categoryReff)
//	if tx.RowsAffected == 0 {
//		status, _ := utils.QueryResponse(false)
//		return &response.Response{Status: status, Message: "INVALID_CATEGORY"}, nil
//	}
//
//	productModel.ID = product.Id
//	productModel.Name = product.Name
//	productModel.Weight = uint(product.Weight)
//	productModel.Category = categoryReff
//	productModel.CategoryId = categoryReff.ID
//	productModel.Price = product.Price
//	productModel.PerUnit = uint(product.PerUnit)
//	productModel.Description = product.Description
//	productModel.ImageUrl = product.ImageUrl
//
//	save := server.DB.Save(&productModel)
//	status, message := utils.ModifyingResponse(save.Error == nil)
//	return &response.Response{
//		Status:  status,
//		Message: message,
//	}, nil
//}
//
//func (server ProductServiceServerImpl) Delete(ctx context.Context, id *ProductId) (*response.Response, error) {
//	server.RedisClient.Del(ctx, id.Id)
//	tx := server.DB.Delete(&model.Product{ID: id.Id})
//	status, message := utils.ModifyingResponse(tx.Error == nil)
//	return &response.Response{Status: status, Message: message}, nil
//}
//
//func (server ProductServiceServerImpl) mustEmbedUnimplementedProductServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}
