package category

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/category"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
	"reflect"
)

type CategoryServiceServerImpl struct {
	Repository  category.CategoryRepositoryImpl
	RedisClient *redis.Client
}

func NewCategoryServiceServerImpl(db *sql.DB) *CategoryServiceServerImpl {
	return &CategoryServiceServerImpl{Repository: category.CategoryRepositoryImpl{DB: db}, RedisClient: utils.NewRedisClient()}
}

func (server *CategoryServiceServerImpl) FindById(context context.Context, categoryId *CategoryId) (*CategoryResponse, error) {
	tx, _ := server.Repository.DB.Begin()
	categoryById := server.Repository.FindById(context, tx, categoryId.String())
	sts, message := utils.FetchResponseForQuerying(!reflect.ValueOf(categoryById).IsZero())
	grpcCategory := Category{
		Category: categoryById.Category,
		Id:       categoryById.Id,
		ImageUrl: categoryById.ImageUrl.(string),
	}
	return &CategoryResponse{
		Data:    &grpcCategory,
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) FindAll(context context.Context, _ *EmptyCategory) (*MultipleCategoryResponse, error) {
	tx, _ := server.Repository.DB.Begin()
	rows := server.Repository.FindAll(context, tx)
	categories := fetchCategories(rows)
	status, message := utils.FetchResponseForQuerying(len(categories) > 0)
	return &MultipleCategoryResponse{
		Status:  status,
		Message: message,
		Data:    categories,
	}, nil
}

func fetchCategories(rows *sql.Rows) []*Category {
	var categoriesModel []*Category
	for rows.Next() {
		cm := Category{}
		var imageUrl sql.NullString
		err := rows.Scan(&cm.Id, &cm.Category, &imageUrl)
		if err != nil {
			panic("scan error")
		}
		if imageUrl.Valid {
			cm.ImageUrl = imageUrl.String
		}
		categoriesModel = append(categoriesModel, &cm)

	}
	return categoriesModel
}

func (server *CategoryServiceServerImpl) Save(context context.Context, category *Category) (*response.Response, error) {
	id, _ := uuid.NewUUID()
	categoryModel := models.CategoryModel{
		Category: category.Category,
		Id:       id.String(),
		Deleted:  false,
	}

	tx, _ := server.Repository.DB.Begin()
	saved := server.Repository.Save(context, tx, categoryModel)
	sts, message := utils.FetchResponseForQuerying(saved)
	return &response.Response{
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) Update(context context.Context, category *Category) (*response.Response, error) {
	categoryModel := models.CategoryModel{
		Category: category.Category,
		ImageUrl: category.ImageUrl,
	}

	tx, _ := server.Repository.DB.Begin()
	updated := server.Repository.Update(context, tx, categoryModel, category.Id)
	sts, message := utils.FetchResponseForQuerying(updated)
	return &response.Response{
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) Delete(context context.Context, categoryId *CategoryId) (*response.Response, error) {
	tx, _ := server.Repository.DB.Begin()
	deleted := server.Repository.Delete(context, tx, categoryId.String())
	sts, message := utils.FetchResponseForModifying(deleted)
	return &response.Response{Status: sts, Message: message}, nil
}
func (server *CategoryServiceServerImpl) mustEmbedUnimplementedCategoryServiceServer() {}
