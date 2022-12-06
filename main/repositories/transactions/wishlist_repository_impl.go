package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type WishlistRepositoryImpl struct {
	DB *sql.DB
}

func (w WishlistRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows {
	query := "SELECT id, name, price, weight, category, per_unit, image_url FROM liked WHERE customer_id = $1"
	rows, err := tx.QueryContext(context, query, userId)
	helpers.PanicIfError(err)
	return rows
}

func (w WishlistRepositoryImpl) FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) models.WishlistModel {
	query := "SELECT id, name, price, weight, category, per_unit, image_url FROM liked WHERE user_id = $1 AND product_id = $2"
	rows := tx.QueryRowContext(context, query, userId, productId)

	wishlist := models.WishlistModel{}
	err := rows.Scan(&wishlist.Id, &wishlist.Name, &wishlist.Price, &wishlist.Weight, &wishlist.Category,
		&wishlist.PerUnit, &wishlist.ImageUrl)
	helpers.PanicIfError(err)
	return wishlist
}

func (w WishlistRepositoryImpl) Save(context context.Context, tx *sql.Tx, product models.WishlistModel) bool {
	sql := "INSERT INTO liked(id, name, category, image_url, per_unit, price, weight, product_id, user_id) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	result, err := tx.ExecContext(context, sql, product.Id, product.Name, product.Category, product.ImageUrl, product.PerUnit,
		product.Price, product.Weight, product.ProductId, product.UserId)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (w WishlistRepositoryImpl) Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool {
	sql := "DELETE FROM liked WHERE user_id = $1 AND product_id = $2"
	result, err := tx.ExecContext(context, sql, userId, productId)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return affected > 0
}
