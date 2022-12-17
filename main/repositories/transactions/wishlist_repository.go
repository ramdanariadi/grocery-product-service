package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type WishlistRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows
	FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *models.WishlistModel
	Save(context context.Context, tx *sql.Tx, product *models.WishlistModel) error
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) error
}
