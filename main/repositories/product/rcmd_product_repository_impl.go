package product

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type RcmdProductRepositoryImpl struct {
	DB *sql.DB
}

func (repository RcmdProductRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel {
	query := "SELECT recommendation_products.product_id as id, name, price, per_unit, weight, category, description, recommendation_products.image_url  " +
		"FROM recommendation_products " +
		"WHERE recommendation_products.product_id = $1"
	rows := tx.QueryRowContext(context, query, id)
	product := models.ProductModel{}
	rows.Scan(&product.Id, &product.Name, &product.Price, &product.PerUnit, &product.Weight, &product.Category, &product.Description,
		&product.ImageUrl)
	return product
}

func (repository RcmdProductRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) []models.ProductModel {
	query := "SELECT recommendation_products.product_id as id, name, price, per_unit, weight, category, description, recommendation_products.image_url  " +
		"FROM recommendation_products"

	rows, err := tx.QueryContext(context, query)
	helpers.PanicIfError(err)
	var topProducts []models.ProductModel
	for rows.Next() {
		productTmp := models.ProductModel{}
		err = rows.Scan(&productTmp.Id, &productTmp.Name, &productTmp.Price, &productTmp.PerUnit,
			&productTmp.Weight, &productTmp.Category, &productTmp.Description, &productTmp.ImageUrl)
		helpers.PanicIfError(err)
		topProducts = append(topProducts, productTmp)
	}
	return topProducts
}

func (repository RcmdProductRepositoryImpl) Save(context context.Context, tx *sql.Tx, saveRequest requestBody.RcmdProductSaveRequest) bool {
	sql := "INSERT INTO recommendation_products(product_id, name, weight, price, per_unit, category, description, image_url, deleted) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	result, err := tx.ExecContext(context, sql, saveRequest.ProductId, saveRequest.Name, saveRequest.Weight, saveRequest.Price, saveRequest.PerUnit, saveRequest.Category, saveRequest.Description, saveRequest.ImageUrl, false)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository RcmdProductRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sql := "DELETE FROM recommendation_products WHERE product_id = $1"
	result, err := tx.ExecContext(context, sql, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}
