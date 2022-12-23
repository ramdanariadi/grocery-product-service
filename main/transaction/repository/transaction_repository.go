package repository

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/model"
)

type TransactionRepository interface {
	FindByTransactionId(context context.Context, tx *sql.Tx, id string) *model.TransactionModel
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []*model.TransactionModel
	Save(context context.Context, tx *sql.Tx, model *model.TransactionModel) error
	Delete(context context.Context, tx *sql.Tx, id string) error
}
