package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
)

type CartRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []models.CartModel
	Save(context context.Context, tx *sql.Tx, product models.CartModel) bool
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool
}
