package category

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func (repository CategoryRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) models.CategoryModel {
	query := "select id, category, image_url " +
		"from category where deleted is false and id = $1"
	row := tx.QueryRowContext(context, query, id)
	cm := models.CategoryModel{}

	err := row.Scan(&cm.Id, &cm.Category, &cm.ImageUrl)
	if err != nil {
		return models.CategoryModel{}
	}
	cm.Deleted = false

	return cm
}

func (repository CategoryRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) []models.CategoryModel {
	query := "select id, category, image_url from category where deleted is false"
	result, err := tx.QueryContext(context, query)
	if err != nil {
		panic("query error")
	}
	var categoriesModel []models.CategoryModel
	for result.Next() {
		cm := models.CategoryModel{}

		err := result.Scan(&cm.Id, &cm.Category, &cm.ImageUrl)
		if err != nil {
			panic("scan error")
		}
		cm.Deleted = false

		categoriesModel = append(categoriesModel, cm)

	}
	return categoriesModel
}

func (repository CategoryRepositoryImpl) Save(context context.Context, tx *sql.Tx, requestBody requestBody.CategorySaveRequest) bool {
	id, _ := uuid.NewUUID()
	fmt.Println(requestBody.Category)
	fmt.Println(id)
	fmt.Println("Image url ", requestBody.ImageUrl)
	sqlInsert := "INSERT INTO category (id, category, image_url, deleted) values($1,$2,$3,$4)"
	result, err := tx.ExecContext(context, sqlInsert, id.String(), requestBody.Category, requestBody.ImageUrl, false)
	//result, err := tx.ExecContext(context, sqlInsert, id.String(), requestBody.Category, requestBody.ImageUrl, false)
	helpers.PanicIfError(err)

	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	fmt.Println("Row affected : ", affected)

	return affected > 0
}

func (repository CategoryRepositoryImpl) Update(context context.Context, tx *sql.Tx, requestBody requestBody.CategorySaveRequest, id string) bool {
	sql := "UPDATE category SET category = $1, image_url = $2 WHERE id = $3"
	result, err := tx.ExecContext(context, sql, requestBody.Category, requestBody.ImageUrl, id)
	helpers.PanicIfError(err)

	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository CategoryRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sql := "DELETE FROM category WHERE id = $1"
	result, err := tx.ExecContext(context, sql, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	return affected > 0
}
