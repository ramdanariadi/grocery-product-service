package transactions

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type CartRepositoryImpl struct {
	DB *sql.DB
}

func (repository CartRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows {
	query := "SELECT id, name, price, weight, category, total, per_unit, image_url FROM cart WHERE user_id = $1"
	rows, err := tx.QueryContext(context, query, userId)
	helpers.PanicIfError(err)
	return rows
}

func (repository CartRepositoryImpl) FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *sql.Row {
	query := "SELECT id, name, price, weight, category, total, per_unit, image_url FROM cart WHERE user_id = $1"
	row := tx.QueryRowContext(context, query, userId)
	return row
}

func (repository CartRepositoryImpl) Save(context context.Context, tx *sql.Tx, product models.CartModel) bool {
	sql := "INSERT INTO cart(id, name, category, total, image_url, per_unit, price, weight, product_id, user_id) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
	id, _ := uuid.NewUUID()
	result, err := tx.ExecContext(context, sql, id, product.Name, product.Category, product.Total,
		product.ImageUrl, product.PerUnit, product.Price, product.Weight, product.ProductId, product.UserId)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository CartRepositoryImpl) Delete(context context.Context, tx *sql.Tx, userId string, wishlistId string) bool {
	sql := "DELETE FROM cart WHERE user_id = $1 AND id = $2"
	result, err := tx.ExecContext(context, sql, userId, wishlistId)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}
