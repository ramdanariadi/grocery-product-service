package category

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/category/model"
	"github.com/ramdanariadi/grocery-product-service/main/category/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
	"log"
	"time"
)

type CategoryServiceServerImpl struct {
	Repository  repository.CategoryRepository
	RedisClient *redis.Client
	DB          *sql.DB
}

func NewCategoryServiceServerImpl(db *sql.DB) *CategoryServiceServerImpl {
	return &CategoryServiceServerImpl{
		Repository:  repository.CategoryRepositoryImpl{},
		RedisClient: setup.NewRedisClient(),
		DB:          db}
}

func (server *CategoryServiceServerImpl) FindById(context context.Context, categoryId *CategoryId) (*CategoryResponse, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	var categoryById *model.CategoryModel
	cache, _ := server.RedisClient.Get(context, categoryId.Id).Result()
	if cache != "" {
		log.Printf("found in cache : %s", cache)
		err := json.Unmarshal([]byte(cache), &categoryById)
		if err != nil {
			status, message := setup.ResponseForQuerying(false)
			return &CategoryResponse{
				Data:    nil,
				Status:  status,
				Message: message,
			}, nil
		}
	} else {
		categoryById = server.Repository.FindById(context, tx, categoryId.Id)
		if categoryById == nil {
			status, message := setup.ResponseForQuerying(false)
			return &CategoryResponse{
				Data:    nil,
				Status:  status,
				Message: message,
			}, nil
		}
		bytes, err := json.Marshal(categoryById)
		utils.PanicIfError(err)
		server.RedisClient.Set(context, categoryId.Id, bytes, 2*time.Hour)
	}

	grpcCategory := Category{
		Category: categoryById.Category,
		Id:       categoryById.Id,
		ImageUrl: categoryById.ImageUrl,
	}
	status, message := setup.ResponseForQuerying(categoryById != nil)
	return &CategoryResponse{
		Data:    &grpcCategory,
		Status:  status,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) FindAll(context context.Context, _ *EmptyCategory) (*MultipleCategoryResponse, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	rows := server.Repository.FindAll(context, tx)
	categories := fetchCategories(rows)
	status, message := setup.ResponseForQuerying(true)
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
	utils.LogIfError(rows.Close())
	return categoriesModel
}

func (server *CategoryServiceServerImpl) Save(context context.Context, category *Category) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	id, _ := uuid.NewUUID()
	categoryModel := model.CategoryModel{
		Category: category.Category,
		Id:       id.String(),
	}

	err = server.Repository.Save(context, tx, &categoryModel)
	sts, message := setup.ResponseForQuerying(err == nil)
	return &response.Response{
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) Update(context context.Context, category *Category) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	categoryModel := model.CategoryModel{
		Category: category.Category,
		ImageUrl: category.ImageUrl,
	}

	err = server.Repository.Update(context, tx, &categoryModel, category.Id)
	_, err = server.RedisClient.Del(context, category.Id).Result()
	utils.LogIfError(err)
	sts, message := setup.ResponseForQuerying(err == nil)
	return &response.Response{
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) Delete(context context.Context, categoryId *CategoryId) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	err = server.Repository.Delete(context, tx, categoryId.Id)
	_, err = server.RedisClient.Del(context, categoryId.Id).Result()
	utils.LogIfError(err)
	sts, message := setup.ResponseForModifying(err == nil)
	return &response.Response{Status: sts, Message: message}, nil
}
func (server *CategoryServiceServerImpl) mustEmbedUnimplementedCategoryServiceServer() {}
