package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type TransactionRepository interface {
	FindByTransactionId(context context.Context, tx *sql.Tx, id string) (*sql.Row, *sql.Rows)
	FindByUserId(context context.Context, tx *sql.Tx, userId string) (*sql.Rows, *sql.Rows)
	Save(context context.Context, tx *sql.Tx, model models.TransactionModel)
	Delete(context context.Context, tx *sql.Tx, id string)
}
