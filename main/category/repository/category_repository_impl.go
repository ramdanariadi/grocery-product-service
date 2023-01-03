package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/category/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"log"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) *model.CategoryModel {
	query := "SELECT id, category, image_url " +
		"FROM categories WHERE deleted_at IS NULL and id = $1"
	row := tx.QueryRowContext(context, query, id)
	cm := model.CategoryModel{}
	var imageUrl sql.NullString
	err := row.Scan(&cm.Id, &cm.Category, &imageUrl)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if imageUrl.Valid {
		cm.ImageUrl = imageUrl.String
	}

	return &cm
}

func (repository CategoryRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) *sql.Rows {
	query := "SELECT id, category, image_url FROM categories WHERE deleted_at IS NULL"
	result, err := tx.QueryContext(context, query)
	utils.LogIfError(err)
	return result
}

func (repository CategoryRepositoryImpl) Save(context context.Context, tx *sql.Tx, requestBody *model.CategoryModel) error {
	id, _ := uuid.NewUUID()
	sqlInsert := "INSERT INTO categories (id, category, image_url, created_at) VALUES($1,$2,$3,NOW())"
	_, err := tx.ExecContext(context, sqlInsert, id, requestBody.Category, requestBody.ImageUrl)
	utils.LogIfError(err)
	return err
}

func (repository CategoryRepositoryImpl) Update(context context.Context, tx *sql.Tx, request *model.CategoryModel, id string) error {
	sql := "UPDATE categories SET category = $1, image_url = $2, updated_at = NOW() WHERE id = $3"
	_, err := tx.ExecContext(context, sql, request.Category, request.ImageUrl, id)
	utils.LogIfError(err)
	return err
}

func (repository CategoryRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) error {
	sql := "UPDATE categories SET deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(context, sql, id)
	utils.LogIfError(err)
	return err
}
