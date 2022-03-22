package transactions

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go-tunas/customresponses/transaction"
	"go-tunas/helpers"
	"go-tunas/models"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func (repository TransactionRepositoryImpl) FindByTransactionId(context context.Context, tx *sql.Tx, id string) transaction.TransactionCustomResponse {
	queryTransaction := "SELECT id, total_price, transaction_date, user_name, user_mobile " +
		"user_email " +
		"FROM transactions " +
		"WHERE id = $1"
	transact := transaction.TransactionCustomResponse{}
	rows, err := tx.QueryContext(context, queryTransaction, id)
	helpers.PanicIfError(err)

	err = rows.Scan(&transact.Id, &transact.TotalPrice, &transact.TransactionDate, &transact.UserName, &transact.UserMobile, &transact.UserEmail)
	helpers.PanicIfError(err)

	queryDetailTransaction := "SELECT name, id, image_url, product_id, price, weight, per_unit," +
		"total, transaction_id " +
		"FROM detail_transaction" +
		"WHERE transaction_id = $1"
	dtRows, err := tx.QueryContext(context, queryDetailTransaction, id)
	helpers.PanicIfError(err)

	for dtRows.Next() {
		product := transaction.ProductTransaction{}
		dtRows.Scan(&product.Name, &product.Id, &product.ImageUrl, &product.ProductId, &product.Price,
			&product.Weight, &product.PerUnit, &product.Total, &product.TransactionId)
		transact.DetailTransaction = append(transact.DetailTransaction, product)
	}
	return transact
}

func (repository TransactionRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) []transaction.TransactionCustomResponse {
	sqlTransaction := "SELECT id, total_price, transaction_date, user_name, user_mobile, user_email, detail_transaction " +
		"FROM transaction WHERE user_id = $1"
	rows, err := tx.QueryContext(context, sqlTransaction, userId)
	helpers.PanicIfError(err)

	sqlDetailTransaction := "SELECT user_name, id, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM detail_transaction dt " +
		"JOIN transaction ON t t.id = dt.transaction_id " +
		"WHERE t.user_id = $1"
	rowDetailTransaction, err := tx.QueryContext(context, sqlDetailTransaction, userId)
	helpers.PanicIfError(err)

	var customTransaction []transaction.TransactionCustomResponse
	for rows.Next() {
		transactionTmp := transaction.TransactionCustomResponse{}
		err := rows.Scan(&transactionTmp.Id, &transactionTmp.TotalPrice, &transactionTmp.TransactionDate, &transactionTmp.UserName,
			&transactionTmp.UserMobile, &transactionTmp.UserEmail, &transactionTmp.DetailTransaction)
		helpers.PanicIfError(err)
		customTransaction = append(customTransaction, transactionTmp)
	}

	for rowDetailTransaction.Next() {
		detailProduct := transaction.ProductTransaction{}
		err := rowDetailTransaction.Scan(&detailProduct.Name, &detailProduct.Id, &detailProduct.ImageUrl, &detailProduct.ProductId,
			&detailProduct.Price, &detailProduct.Weight, &detailProduct.PerUnit, &detailProduct.Total, &detailProduct.TransactionId)
		helpers.PanicIfError(err)

		for _, tran := range customTransaction {
			if tran.Id == detailProduct.Id {
				tran.DetailTransaction = append(tran.DetailTransaction, detailProduct)
			}
		}
	}

	return customTransaction
}

func (repository TransactionRepositoryImpl) Save(context context.Context, tx *sql.Tx, model models.TransactionModel) bool {
	sqlTransaction := "INSERT INTO transaction(total_price,transaction_date,user_id,id,user_email,user_mobile,user_name) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7)"
	sqlDetailTransaction := "INSERT INTO detail_transaction(transaction_id,product_id,id,per_unit,price,total,weight,image_url,name) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	result, err := tx.ExecContext(context, sqlTransaction, model.Total, "date", model.UserId, model.Id)
	helpers.PanicIfError(err)

	for _, dt := range model.DetailTransaction {
		id, err := uuid.NewUUID()
		helpers.PanicIfError(err)

		result, err := tx.ExecContext(context, sqlDetailTransaction, model.Id, "productId", id, dt.PerUnit, dt.Price, "total", dt.Weight, dt.ImageUrl, dt.Name)
		helpers.PanicIfError(err)

		affected, err := result.RowsAffected()
		helpers.PanicIfError(err)

		if affected > 0 {
			panic("err add detail transaction")
		}
	}

	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	return affected > 0
}

func (repository TransactionRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sqlDetailTransaction := "DELETE FROM detail_transaction WHERE id_transaction = $1"
	result, err := tx.ExecContext(context, sqlDetailTransaction, id)
	helpers.PanicIfError(err)

	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	if affected == 0 {
		return false
	}

	sqlTransaction := "DELETE FROM transaction WHERE id = $1"
	execContext, err := tx.ExecContext(context, sqlTransaction, id)
	helpers.PanicIfError(err)

	rowsAffected, err := execContext.RowsAffected()
	helpers.PanicIfError(err)

	return rowsAffected > 0
}
