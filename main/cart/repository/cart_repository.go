package repository

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/cart/model"
)

type CartRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows
	FindByUserAndProductId(context context.Context, tx *sql.Tx, userId string, productId string) *sql.Row
	Save(context context.Context, tx *sql.Tx, product *model.CartModel) error
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) error
}
