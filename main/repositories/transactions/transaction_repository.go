package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/transaction"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
)

type TransactionRepository interface {
	FindByTransactionId(context context.Context, tx *sql.Tx, id string) transaction.TransactionCustomResponse
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []transaction.TransactionCustomResponse
	Save(context context.Context, tx *sql.Tx, model models.TransactionModel) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}
