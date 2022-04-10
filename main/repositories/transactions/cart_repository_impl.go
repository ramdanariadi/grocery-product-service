package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
)

type CartRepositoryImpl struct {
	DB *sql.DB
}

func (repository CartRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) []models.CartModel {
	query := "SELECT id, name, price, weight, category, total, per_unit, image_url FROM cart WHERE user_id = $1"
	rows, err := tx.QueryContext(context, query, userId)
	helpers.PanicIfError(err)

	var carts []models.CartModel

	for rows.Next() {
		cart := models.CartModel{}
		err := rows.Scan(&cart.Id, &cart.Name, &cart.Price, &cart.Weight, &cart.Category,
			&cart.Total, &cart.PerUnit, &cart.ImageUrl)
		helpers.PanicIfError(err)
		carts = append(carts, cart)
	}
	return carts
}

func (repository CartRepositoryImpl) Save(context context.Context, tx *sql.Tx, product models.CartModel) bool {
	sql := "INSERT INTO cart(id, name, category, total, image_url, per_unit, price, weight, product_id, user_id) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
	result, err := tx.ExecContext(context, sql, product.Id, product.Name, product.Category, product.Total,
		product.ImageUrl, product.PerUnit, product.Price, product.Weight, product.ProductId, product.UserId)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository CartRepositoryImpl) Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool {
	sql := "DELETE FROM cart WHERE user_id = $1 AND product_id = $2"
	result, err := tx.ExecContext(context, sql, userId, productId)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}
