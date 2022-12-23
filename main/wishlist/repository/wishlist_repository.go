package repository

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
)

type WishlistRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows
	FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *model.WishlistModel
	Save(context context.Context, tx *sql.Tx, product *model.WishlistModel) error
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) error
}
