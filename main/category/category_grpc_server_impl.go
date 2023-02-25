package category

//
//import (
//	"encoding/json"
//	"github.com/go-redis/redis/v8"
//	"github.com/google/uuid"
//	"github.com/ramdanariadi/grocery-product-service/main/category/model"
//	"github.com/ramdanariadi/grocery-product-service/main/response"
//	"github.com/ramdanariadi/grocery-product-service/main/setup"
//	"github.com/ramdanariadi/grocery-product-service/main/utils"
//	"golang.org/x/net/context"
//	"gorm.io/gorm"
//)
//
//type CategoryServiceServerImpl struct {
//	DB          *gorm.DB
//	RedisClient *redis.Client
//}
//
//func NewCategoryServiceServerImpl(db *gorm.DB) *CategoryServiceServerImpl {
//	return &CategoryServiceServerImpl{DB: db, RedisClient: setup.NewRedisClient()}
//}
//
//func (server *CategoryServiceServerImpl) FindById(ctx context.Context, categoryId *CategoryId) (*CategoryResponse, error) {
//	var categoryModel *model.Category
//
//	result, err := server.RedisClient.Get(ctx, categoryId.Id).Result()
//	utils.LogIfError(err)
//
//	if err == nil {
//		err := json.Unmarshal([]byte(result), categoryModel)
//		utils.LogIfError(err)
//	} else {
//		tx := server.DB.First(categoryModel, "id = ?", categoryId.Id)
//		utils.LogIfError(tx.Error)
//		if tx.RowsAffected == 0 {
//			status, message := utils.QueryResponse(false)
//			return &CategoryResponse{
//				Data:    nil,
//				Status:  status,
//				Message: message,
//			}, nil
//		}
//	}
//
//	status, message := utils.QueryResponse(true)
//	grpcCategory := Category{
//		Category: categoryModel.Category,
//		Id:       categoryModel.ID,
//		ImageUrl: categoryModel.ImageUrl,
//	}
//	return &CategoryResponse{
//		Data:    &grpcCategory,
//		Status:  status,
//		Message: message,
//	}, nil
//}
//
//func (server *CategoryServiceServerImpl) FindAll(_ context.Context, _ *EmptyCategory) (*MultipleCategoryResponse, error) {
//	var categoriesModel []*model.Category
//	server.DB.Find(&categoriesModel)
//	categories := fetchCategories(categoriesModel)
//	status, message := utils.QueryResponse(true)
//	return &MultipleCategoryResponse{
//		Status:  status,
//		Message: message,
//		Data:    categories,
//	}, nil
//}
//
//func fetchCategories(categoriesModel []*model.Category) []*Category {
//	var categories []*Category
//	for _, c := range categoriesModel {
//		categories = append(categories, &Category{Id: c.ID, Category: c.Category, ImageUrl: c.ImageUrl})
//	}
//	return categories
//}
//
//func (server *CategoryServiceServerImpl) Save(_ context.Context, category *Category) (*response.Response, error) {
//	id, _ := uuid.NewUUID()
//	save := server.DB.Save(&model.Category{ID: id.String(), Category: category.Category, ImageUrl: category.ImageUrl})
//	sts, message := utils.QueryResponse(save.Error == nil)
//	return &response.Response{
//		Status:  sts,
//		Message: message,
//	}, nil
//}
//
//func (server *CategoryServiceServerImpl) Update(ctx context.Context, category *Category) (*response.Response, error) {
//	categoryModel := model.Category{ID: category.Id}
//	tx := server.DB.First(&categoryModel)
//
//	if tx.RowsAffected == 0 {
//		sts, message := utils.QueryResponse(false)
//		return &response.Response{
//			Status:  sts,
//			Message: message,
//		}, nil
//	}
//
//	server.RedisClient.Del(ctx, category.Id)
//	categoryModel.Category = category.Category
//	categoryModel.ImageUrl = category.ImageUrl
//	save := server.DB.Save(&categoryModel)
//	sts, message := utils.QueryResponse(save.Error == nil)
//	return &response.Response{
//		Status:  sts,
//		Message: message,
//	}, nil
//}
//
//func (server *CategoryServiceServerImpl) Delete(ctx context.Context, categoryId *CategoryId) (*response.Response, error) {
//	server.RedisClient.Del(ctx, categoryId.Id)
//	server.DB.Delete(&model.Category{ID: categoryId.Id})
//	sts, message := utils.ModifyingResponse(true)
//	return &response.Response{Status: sts, Message: message}, nil
//}
//
//func (server *CategoryServiceServerImpl) mustEmbedUnimplementedCategoryServiceServer() {}
