package transactions

import (
	"context"
	"database/sql"
	"go-tunas/models"
)

type CartRepository interface {
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []models.CartModel
	Save(context context.Context, tx *sql.Tx, product models.CartModel) bool
	Delete(context context.Context, tx *sql.Tx, userId string, productId string) bool
}
