package category

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"log"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository CategoryRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) *models.CategoryModel {
	query := "SELECT id, category, image_url " +
		"FROM category WHERE deleted_at IS NULL and id = $1"
	row := tx.QueryRowContext(context, query, id)
	cm := models.CategoryModel{}
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
	query := "SELECT id, category, image_url FROM category WHERE deleted_at IS NULL"
	result, err := tx.QueryContext(context, query)
	helpers.LogIfError(err)
	return result
}

func (repository CategoryRepositoryImpl) Save(context context.Context, tx *sql.Tx, requestBody *models.CategoryModel) error {
	id, _ := uuid.NewUUID()
	sqlInsert := "INSERT INTO category (id, category, image_url, created_at) VALUES($1,$2,$3,NOW())"
	_, err := tx.ExecContext(context, sqlInsert, id, requestBody.Category, requestBody.ImageUrl)
	helpers.LogIfError(err)
	return err
}

func (repository CategoryRepositoryImpl) Update(context context.Context, tx *sql.Tx, request *models.CategoryModel, id string) error {
	sql := "UPDATE category SET category = $1, image_url = $2, updated_at = NOW() WHERE id = $3"
	_, err := tx.ExecContext(context, sql, request.Category, request.ImageUrl, id)
	helpers.LogIfError(err)
	return err
}

func (repository CategoryRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) error {
	sql := "UPDATE category SET deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(context, sql, id)
	helpers.LogIfError(err)
	return err
}
