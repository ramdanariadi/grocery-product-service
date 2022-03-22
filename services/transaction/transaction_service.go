package transaction

import (
	"go-tunas/customresponses/transaction"
	"go-tunas/models"
)

type TransactionService interface {
	FindByTransactionId(id string) transaction.TransactionCustomResponse
	FindByUserId(id string) []transaction.TransactionCustomResponse
	Save(transaction models.TransactionModel) bool
	Delete(id string) bool
}
