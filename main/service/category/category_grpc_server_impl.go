package category

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/category"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
)

type CategoryServiceServerImpl struct {
	Repository  category.CategoryRepositoryImpl
	RedisClient *redis.Client
}

func NewCategoryServiceServerImpl(db *sql.DB) *CategoryServiceServerImpl {
	return &CategoryServiceServerImpl{Repository: category.CategoryRepositoryImpl{DB: db}, RedisClient: utils.NewRedisClient()}
}

func (server *CategoryServiceServerImpl) FindById(context context.Context, categoryId *CategoryId) (*CategoryResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	categoryById := server.Repository.FindById(context, tx, categoryId.Id)
	status, message := utils.ResponseForQuerying(!utils.IsStructEmpty(categoryById))
	grpcCategory := Category{
		Category: categoryById.Category,
		Id:       categoryById.Id,
		ImageUrl: categoryById.ImageUrl,
	}
	return &CategoryResponse{
		Data:    &grpcCategory,
		Status:  status,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) FindAll(context context.Context, _ *EmptyCategory) (*MultipleCategoryResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	rows := server.Repository.FindAll(context, tx)
	categories := fetchCategories(rows)
	status, message := utils.ResponseForQuerying(len(categories) > 0)
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
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	id, _ := uuid.NewUUID()
	categoryModel := models.CategoryModel{
		Category: category.Category,
		Id:       id.String(),
		Deleted:  false,
	}

	saved := server.Repository.Save(context, tx, categoryModel)
	sts, message := utils.ResponseForQuerying(saved)
	return &response.Response{
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) Update(context context.Context, category *Category) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categoryModel := models.CategoryModel{
		Category: category.Category,
		ImageUrl: category.ImageUrl,
	}

	updated := server.Repository.Update(context, tx, categoryModel, category.Id)
	sts, message := utils.ResponseForQuerying(updated)
	return &response.Response{
		Status:  sts,
		Message: message,
	}, nil
}
func (server *CategoryServiceServerImpl) Delete(context context.Context, categoryId *CategoryId) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	deleted := server.Repository.Delete(context, tx, categoryId.Id)
	sts, message := utils.ResponseForModifying(deleted)
	return &response.Response{Status: sts, Message: message}, nil
}
func (server *CategoryServiceServerImpl) mustEmbedUnimplementedCategoryServiceServer() {}
