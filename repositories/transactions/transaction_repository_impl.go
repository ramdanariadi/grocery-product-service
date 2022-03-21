package transactions

import (
	"context"
	"database/sql"
	"go-tunas/customresponses/transaction"
	"go-tunas/models"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func (repository TransactionRepositoryImpl) FindByTransactionId(context context.Context, tx *sql.Tx, id string) transaction.TransactionCustomResponse {
	//TODO implement me
	panic("implement me")
}

func (repository TransactionRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) []transaction.TransactionCustomResponse {
	//TODO implement me
	panic("implement me")
}

func (repository TransactionRepositoryImpl) Save(context context.Context, tx *sql.Tx, model models.TransactionModel) bool {
	//TODO implement me
	panic("implement me")
}

func (repository TransactionRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	//TODO implement me
	panic("implement me")
}
