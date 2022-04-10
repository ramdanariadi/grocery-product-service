package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/product"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
)

type WishlistRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []product.WishlistResponse
	FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) product.WishlistResponse
	Save(context context.Context, tx *sql.Tx, product models.WishlistModel) bool
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool
}
