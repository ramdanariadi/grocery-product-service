package product

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func (repository ProductRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel {
	query := "SELECT products.id, name, price, per_unit, weight, category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN category ON products.category_id = category.id " +
		"WHERE products.id = $1"
	row := tx.QueryRowContext(context, query, id)
	product := models.ProductModel{}
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.PerUnit, &product.Weight,
		&product.Category, &product.CategoryId,
		&product.Description, &product.ImageUrl)
	if err != nil {
		return models.ProductModel{}
	}
	//helpers.PanicIfError(err)
	return product
}

func (repository ProductRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN category ON products.category_id = category.id"

	rows, err := tx.QueryContext(context, query)
	helpers.PanicIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) FindByCategory(context context.Context, tx *sql.Tx, id string) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN category ON products.category_id = category.id " +
		"WHERE products.category_id = $1"
	rows, err := tx.QueryContext(context, query, id)
	helpers.PanicIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) Save(context context.Context, tx *sql.Tx, product models.ProductModel) bool {
	sql := "INSERT INTO products(id, name, weight, price, per_unit, category_id, description, image_url, deleted, is_top, is_recommended) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	result, err := tx.ExecContext(context, sql, product.Id, product.Name, product.Weight, product.Price,
		product.PerUnit, product.CategoryId, product.Description, product.ImageUrl, false, false, false)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository ProductRepositoryImpl) Update(context context.Context, tx *sql.Tx, product models.ProductModel) bool {
	sql := "UPDATE products SET name=$1, price=$2, weight=$3, category_id=$4, per_unit=$5," +
		"description=$6, image_url=$7, is_top=$8, is_recommended=$9" +
		"WHERE id = $10"
	result, err := tx.ExecContext(context, sql, product.Name, product.Price, product.Weight,
		product.Category, product.PerUnit, product.Description, product.ImageUrl, product.IsTop,
		product.IsRecommended, product.Id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	return affected > 0
}

func (repository ProductRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sql := "UPDATE products set deleted = true WHERE id = $1"
	result, err := tx.ExecContext(context, sql, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}
