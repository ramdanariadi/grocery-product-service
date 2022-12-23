package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
	"log"
)

type WishlistRepositoryImpl struct {
	DB *sql.DB
}

func (w WishlistRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows {
	query := "SELECT id, name, price, weight, category, per_unit, image_url FROM liked WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := tx.QueryContext(context, query, userId)
	helpers.LogIfError(err)
	return rows
}

func (w WishlistRepositoryImpl) FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *model.WishlistModel {
	query := "SELECT id, name, price, weight, category, per_unit, image_url FROM liked WHERE user_id = $1 AND product_id = $2 AND deleted_at IS NULL"
	rows := tx.QueryRowContext(context, query, userId, productId)

	wishlist := model.WishlistModel{}
	err := rows.Scan(&wishlist.Id, &wishlist.Name, &wishlist.Price, &wishlist.Weight, &wishlist.Category,
		&wishlist.PerUnit, &wishlist.ImageUrl)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &wishlist
}

func (w WishlistRepositoryImpl) Save(context context.Context, tx *sql.Tx, product *model.WishlistModel) error {
	sql := "INSERT INTO liked(id, name, category, image_url, per_unit, price, weight, product_id, user_id, created_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW())"
	id, _ := uuid.NewUUID()
	_, err := tx.ExecContext(context, sql, id, product.Name, product.Category, product.ImageUrl, product.PerUnit,
		product.Price, product.Weight, product.ProductId, product.UserId)
	helpers.LogIfError(err)
	return err
}

func (w WishlistRepositoryImpl) Delete(context context.Context, tx *sql.Tx, userId string, wishlistId string) error {
	sql := "UPDATE liked SET deleted_at = NOW() WHERE user_id = $1 AND id = $2"
	_, err := tx.ExecContext(context, sql, userId, wishlistId)
	helpers.LogIfError(err)
	return err
}
