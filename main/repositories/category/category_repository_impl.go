package category

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository CategoryRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) models.CategoryModel {
	query := "SELECT id, category, image_url " +
		"FROM category WHERE deleted_at IS NULL and id = $1"
	row := tx.QueryRowContext(context, query, id)
	cm := models.CategoryModel{}
	var imageUrl sql.NullString
	err := row.Scan(&cm.Id, &cm.Category, &imageUrl)
	if err != nil {
		return models.CategoryModel{}
	}

	if imageUrl.Valid {
		cm.ImageUrl = imageUrl.String
	}

	return cm
}

func (repository CategoryRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) *sql.Rows {
	query := "SELECT id, category, image_url FROM category WHERE deleted_at IS NULL"
	result, err := tx.QueryContext(context, query)
	if err != nil {
		panic("query error")
	}
	return result
}

func (repository CategoryRepositoryImpl) Save(context context.Context, tx *sql.Tx, requestBody models.CategoryModel) bool {
	id, _ := uuid.NewUUID()
	fmt.Println(requestBody.Category)
	fmt.Println(id)
	fmt.Println("Image url ", requestBody.ImageUrl)
	sqlInsert := "INSERT INTO category (id, category, image_url, deleted_at) VALUES($1,$2,$3,$4)"
	result, err := tx.ExecContext(context, sqlInsert, id.String(), requestBody.Category, requestBody.ImageUrl, false)
	helpers.PanicIfError(err)

	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	fmt.Println("Row affected : ", affected)

	return affected > 0
}

func (repository CategoryRepositoryImpl) Update(context context.Context, tx *sql.Tx, request models.CategoryModel, id string) bool {
	sql := "UPDATE category SET category = $1, image_url = $2, updated_at = NOW() WHERE id = $3"
	result, err := tx.ExecContext(context, sql, request.Category, request.ImageUrl, id)
	helpers.PanicIfError(err)

	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository CategoryRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sql := "UPDATE category SET deleted_at = NOW() WHERE id = $1"
	result, err := tx.ExecContext(context, sql, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	return affected > 0
}
