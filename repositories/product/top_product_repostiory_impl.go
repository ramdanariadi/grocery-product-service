package product

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/requestBody"
)

type TopProductRepositoryImpl struct {
	DB *sql.DB
}

func (repository TopProductRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel {
	query := "SELECT top_products.id, name, price, per_unit, weight, category, description, top_products.image_url  " +
		"FROM top_products " +
		"WHERE top_products.id = ?"
	rows, err := tx.QueryContext(context, query, id)
	helpers.PanicIfError(err)
	product := models.ProductModel{}
	err = rows.Scan(product.Id, product.Name, product.Price, product.PerUnit, product.Weight, product.Category, product.Description, product.ImageUrl)
	helpers.PanicIfError(err)
	return product
}

func (repository TopProductRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) []models.ProductModel {
	query := "SELECT top_products.id, name, price, per_unit, weight, category, description, top_products.image_url  " +
		"FROM top_products"

	rows, err := tx.QueryContext(context, query)
	helpers.PanicIfError(err)
	var topProducts []models.ProductModel
	for rows.Next() {
		productTmp := models.ProductModel{}
		err = rows.Scan(productTmp.Id, productTmp.Name, productTmp.Price, productTmp.PerUnit, productTmp.Weight, productTmp.Category, productTmp.Description, productTmp.ImageUrl)
		helpers.PanicIfError(err)
		topProducts = append(topProducts, productTmp)
	}
	return topProducts
}

func (repository TopProductRepositoryImpl) Save(context context.Context, tx *sql.Tx, saveRequest requestBody.TopProductSaveRequest) bool {
	sql := "INSERT INTO top_products(id, name, weight, price, per_unit, category_id, description, image_url, deleted) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	id, _ := uuid.NewUUID()
	result, err := tx.ExecContext(context, sql, id, saveRequest.Name, saveRequest.Weight, saveRequest.Price, saveRequest.PerUnit, saveRequest.Category, saveRequest.Description, saveRequest.ImageUrl, false)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository TopProductRepositoryImpl) Update(context context.Context, tx *sql.Tx, updateRequest requestBody.TopProductSaveRequest, id string) bool {
	sql := "UPDATE FROM top_products SET name=$1, price=$2, weight=$3, category_id=$4, per_unit=$5," +
		"description=$6, image_url=$7" +
		"WHERE id = $8"
	result, err := tx.ExecContext(context, sql, updateRequest.Name, updateRequest.Price, updateRequest.Weight,
		updateRequest.Category, updateRequest.PerUnit, updateRequest.Description, updateRequest.ImageUrl, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	return affected > 0
}

func (repository TopProductRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sql := "DELETE FROM top_products WHERE id = ?"
	result, err := tx.ExecContext(context, sql, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}
