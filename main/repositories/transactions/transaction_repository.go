package transactions

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type TransactionRepository interface {
	FindByTransactionId(context context.Context, tx *sql.Tx, id string) *models.TransactionModel
	FindByUserId(context context.Context, tx *sql.Tx, userId string) []*models.TransactionModel
	Save(context context.Context, tx *sql.Tx, model *models.TransactionModel) error
	Delete(context context.Context, tx *sql.Tx, id string) error
}
