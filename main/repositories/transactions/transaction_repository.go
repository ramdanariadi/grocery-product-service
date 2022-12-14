package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type TransactionRepository interface {
	FindByTransactionId(context context.Context, tx *sql.Tx, id string) *sql.Row
	FindByUserId(context context.Context, tx *sql.Tx, userId string) *sql.Rows
	Save(context context.Context, tx *sql.Tx, model models.TransactionModel) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}
