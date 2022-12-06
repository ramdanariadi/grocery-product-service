package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type CartRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows
	FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *sql.Row
	Save(context context.Context, tx *sql.Tx, product models.CartModel) bool
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool
}
