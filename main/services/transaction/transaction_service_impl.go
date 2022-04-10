package transaction

import (
	"context"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/transaction"
	helpers2 "github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/transactions"
)

type TransactinoServiceImpl struct {
	Repository transactions.TransactionRepositoryImpl
}

func (service TransactinoServiceImpl) FindByTransactionId(id string) transaction.TransactionCustomResponse {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.FindByTransactionId(context.Background(), tx, id)
}

func (service TransactinoServiceImpl) FindByUserId(id string) []transaction.TransactionCustomResponse {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.FindByUserId(context.Background(), tx, id)
}

func (service TransactinoServiceImpl) Save(transaction models.TransactionModel) bool {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Save(context.Background(), tx, transaction)
}

func (service TransactinoServiceImpl) Delete(id string) bool {
	tx, err := service.Repository.DB.Begin()
	helpers2.PanicIfError(err)
	defer helpers2.CommitOrRollback(tx)
	return service.Repository.Delete(context.Background(), tx, id)
}
