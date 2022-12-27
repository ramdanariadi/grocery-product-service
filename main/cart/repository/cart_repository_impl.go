package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/cart/model"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
)

type CartRepositoryImpl struct {
	DB *sql.DB
}

func (repository CartRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows {
	query := "SELECT id, name, price, weight, category, total, per_unit, image_url, product_id FROM cart WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := tx.QueryContext(context, query, userId)
	helpers.LogIfError(err)
	return rows
}

func (repository CartRepositoryImpl) FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *sql.Row {
	query := "SELECT id, name, price, weight, category, total, per_unit, image_url, product_id FROM cart WHERE user_id = $1 AND product_id = $2 AND deleted_at IS NULL"
	row := tx.QueryRowContext(context, query, userId, productId)
	return row
}

func (repository CartRepositoryImpl) Save(context context.Context, tx *sql.Tx, product *model.CartModel) error {
	sql := "INSERT INTO cart(id, name, category, total, image_url, per_unit, price, weight, product_id, user_id, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())"
	id, _ := uuid.NewUUID()
	_, err := tx.ExecContext(context, sql, id, product.Name, product.Category, product.Total,
		product.ImageUrl, product.PerUnit, product.Price, product.Weight, product.ProductId, product.UserId)
	helpers.LogIfError(err)
	return err
}

func (repository CartRepositoryImpl) Update(context context.Context, tx *sql.Tx, cart *model.CartModel) error {
	sql := "UPDATE cart SET name=$2, category=$3, total=$4, image_url=$5, per_unit=$6, price=$7, weight=$8, " +
		"product_id=$9, user_id=$10, updated_at=NOW() " +
		"WHERE id=$1"
	_, err := tx.ExecContext(context, sql, cart.Id, cart.Name, cart.Category, cart.Total,
		cart.ImageUrl, cart.PerUnit, cart.Price, cart.Weight, cart.ProductId, cart.UserId)
	helpers.LogIfError(err)
	return err
}

func (repository CartRepositoryImpl) Delete(context context.Context, tx *sql.Tx, userId string, productId string) error {
	sql := "UPDATE cart SET deleted_at = NOW() WHERE user_id = $1 AND product_id = $2"
	_, err := tx.ExecContext(context, sql, userId, productId)
	helpers.LogIfError(err)
	return err
}
