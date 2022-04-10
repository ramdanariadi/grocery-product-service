package transaction

import (
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/transaction"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
)

type TransactionService interface {
	FindByTransactionId(id string) transaction.TransactionCustomResponse
	FindByUserId(id string) []transaction.TransactionCustomResponse
	Save(transaction models.TransactionModel) bool
	Delete(id string) bool
}
