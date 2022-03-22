package transaction

import (
	"context"
	"go-tunas/customresponses/transaction"
	"go-tunas/helpers"
	"go-tunas/models"
	"go-tunas/repositories/transactions"
)

type TransactinoServiceImpl struct {
	Repository transactions.TransactionRepositoryImpl
}

func (service TransactinoServiceImpl) FindByTransactionId(id string) transaction.TransactionCustomResponse {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.FindByTransactionId(context.Background(), tx, id)
}

func (service TransactinoServiceImpl) FindByUserId(id string) []transaction.TransactionCustomResponse {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.FindByUserId(context.Background(), tx, id)
}

func (service TransactinoServiceImpl) Save(transaction models.TransactionModel) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Save(context.Background(), tx, transaction)
}

func (service TransactinoServiceImpl) Delete(id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	return service.Repository.Delete(context.Background(), tx, id)
}
