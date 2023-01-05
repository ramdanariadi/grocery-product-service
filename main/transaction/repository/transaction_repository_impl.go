package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"log"
	"strings"
	"time"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func (repository TransactionRepositoryImpl) FindByTransactionId(context context.Context, tx *sql.Tx, id string) *model.TransactionModel {
	queryTransaction := "SELECT id, total_price, created_at " +
		"FROM transactions " +
		"WHERE id = $1 AND deleted_at IS NULL"
	row := tx.QueryRowContext(context, queryTransaction, id)
	transactionModel := model.TransactionModel{}
	var transactionDate sql.NullTime
	err := row.Scan(&transactionModel.Id, &transactionModel.TotalPrice, &transactionDate)
	if transactionDate.Valid {
		log.Println("transaction date valid")
		transactionModel.TransactionDate = transactionDate.Time.UnixMilli()
	}
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	queryDetailTransaction := "SELECT id, name, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM transaction_details " +
		"WHERE transaction_id = $1 AND deleted_at IS NULL"
	dtRows, err := tx.QueryContext(context, queryDetailTransaction, id)
	utils.LogIfError(err)

	var detailTransactions []*model.DetailTransactionProductModel
	for dtRows.Next() {
		detailTransaction := model.DetailTransactionProductModel{}
		var imageUrl sql.NullString
		err = dtRows.Scan(&detailTransaction.Id, &detailTransaction.Name, &imageUrl, &detailTransaction.ProductId,
			&detailTransaction.Price, &detailTransaction.Weight, &detailTransaction.PerUnit, &detailTransaction.Total, &detailTransaction.TransactionId)
		if err != nil {
			continue
		}

		if imageUrl.Valid {
			detailTransaction.ImageUrl = imageUrl.String
		}
		detailTransactions = append(detailTransactions, &detailTransaction)
	}
	utils.LogIfError(dtRows.Close())
	attachDetailTransaction(&transactionModel, detailTransactions)
	return &transactionModel
}

func attachDetailTransaction(transaction *model.TransactionModel, detailTransaction []*model.DetailTransactionProductModel) {
	for _, dt := range detailTransaction {
		if dt.TransactionId == transaction.Id {
			transaction.DetailTransaction = append(transaction.DetailTransaction, dt)
		}
	}
}

func (repository TransactionRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) []*model.TransactionModel {
	sqlDetailTransaction := "SELECT dt.id, name, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM transaction_details dt " +
		"JOIN transactions t ON t.id = dt.transaction_id " +
		"WHERE t.user_id = $1 AND dt.deleted_at IS NULL"
	detailTransactionRows, err := tx.QueryContext(context, sqlDetailTransaction, userId)
	utils.PanicIfError(err)
	var detailTransactions []*model.DetailTransactionProductModel
	for detailTransactionRows.Next() {
		detailTransaction := model.DetailTransactionProductModel{}
		var imageUrl sql.NullString
		err = detailTransactionRows.Scan(&detailTransaction.Id, &detailTransaction.Name, &imageUrl, &detailTransaction.ProductId,
			&detailTransaction.Price, &detailTransaction.Weight, &detailTransaction.PerUnit, &detailTransaction.Total, &detailTransaction.TransactionId)
		if err != nil {
			continue
		}

		if imageUrl.Valid {
			detailTransaction.ImageUrl = imageUrl.String
		}
		detailTransactions = append(detailTransactions, &detailTransaction)
	}
	utils.LogIfError(detailTransactionRows.Close())

	sqlTransaction := "SELECT id, total_price, created_at " +
		"FROM transactions WHERE user_id = $1 AND deleted_at IS NULL"
	transactionRows, err := tx.QueryContext(context, sqlTransaction, userId)
	utils.PanicIfError(err)

	var transactions []*model.TransactionModel
	for transactionRows.Next() {
		transactionModel := model.TransactionModel{}
		var transactionDate sql.NullTime
		err = transactionRows.Scan(&transactionModel.Id, &transactionModel.TotalPrice, &transactionDate)
		if err != nil {
			continue
		}

		if transactionDate.Valid {
			transactionModel.TransactionDate = transactionDate.Time.UnixMilli()
		}

		attachDetailTransaction(&transactionModel, detailTransactions)
		transactions = append(transactions, &transactionModel)
	}
	utils.LogIfError(transactionRows.Close())
	return transactions
}

func (repository TransactionRepositoryImpl) Save(context context.Context, tx *sql.Tx, model *model.TransactionModel) error {
	sqlTransaction := "INSERT INTO transactions(id, total_price, user_id, created_at) VALUES($1,$2,$3, NOW())"
	transactionId, _ := uuid.NewUUID()
	_, err := tx.ExecContext(context, sqlTransaction, transactionId, model.TotalPrice, model.UserId)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var statement []string
	var values []interface{}
	for index, dt := range model.DetailTransaction {
		statement = append(statement, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			index*9+1,
			index*9+2,
			index*9+3,
			index*9+4,
			index*9+5,
			index*9+6,
			index*9+7,
			index*9+8,
			index*9+9,
			index*9+10))

		id, err := uuid.NewUUID()
		utils.PanicIfError(err)
		now := time.Now()
		values = append(values, transactionId, dt.ProductId, id, dt.PerUnit, dt.Price, dt.Total, dt.Weight, dt.ImageUrl, dt.Name, now.Format("2006-01-02 15:04:05"))
	}

	sqlDetailTransaction := fmt.Sprintf("INSERT INTO transaction_details(transaction_id,product_id,id,per_unit,price,total,weight,image_url,name, created_at) "+
		"VALUES %s", strings.Join(statement, ","))

	_, err = tx.ExecContext(context, sqlDetailTransaction, values...)
	utils.LogIfError(err)
	return err
}

func (repository TransactionRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) error {
	sqlDetailTransaction := "UPDATE transaction_details SET deleted_at = NOW() WHERE transaction_id = $1"
	_, err := tx.ExecContext(context, sqlDetailTransaction, id)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	sqlTransaction := "UPDATE transactions SET deleted_at = NOW() WHERE id = $1"
	_, err = tx.ExecContext(context, sqlTransaction, id)
	utils.LogIfError(err)
	return err
}
