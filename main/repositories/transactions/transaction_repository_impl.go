package transactions

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses/transaction"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"log"
	"strings"
	"time"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func (repository TransactionRepositoryImpl) FindByTransactionId(context context.Context, tx *sql.Tx, id string) transaction.TransactionCustomResponse {
	queryTransaction := "SELECT id, total_price, transaction_date, user_name, user_mobile " +
		"user_email " +
		"FROM transaction " +
		"WHERE id = $1"
	transact := transaction.TransactionCustomResponse{}
	rows := tx.QueryRowContext(context, queryTransaction, id)

	err := rows.Scan(&transact.Id, &transact.TotalPrice, &transact.TransactionDate, &transact.UserName, &transact.UserMobile, &transact.UserEmail)
	if err != nil {
		return transact
	}

	queryDetailTransaction := "SELECT name, id, image_url, product_id, price, weight, per_unit," +
		"total, transaction_id " +
		"FROM detail_transaction " +
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
	sqlTransaction := "SELECT id, total_price, transaction_date, user_name, user_mobile, user_email " +
		"FROM transaction WHERE user_id = $1"
	rows, err := tx.QueryContext(context, sqlTransaction, userId)
	helpers.PanicIfError(err)

	var customTransaction []transaction.TransactionCustomResponse
	for rows.Next() {
		transactionTmp := transaction.TransactionCustomResponse{}
		err := rows.Scan(&transactionTmp.Id, &transactionTmp.TotalPrice, &transactionTmp.TransactionDate, &transactionTmp.UserName,
			&transactionTmp.UserMobile, &transactionTmp.UserEmail)
		helpers.PanicIfError(err)
		customTransaction = append(customTransaction, transactionTmp)
	}
	rows.Close()

	sqlDetailTransaction := "SELECT name, dt.id, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM detail_transaction dt " +
		"JOIN transaction t ON t.id = dt.transaction_id " +
		"WHERE t.user_id = $1"
	rowDetailTransaction, err := tx.QueryContext(context, sqlDetailTransaction, userId)
	helpers.PanicIfError(err)

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
	rowDetailTransaction.Close()

	return customTransaction
}

func (repository TransactionRepositoryImpl) Save(context context.Context, tx *sql.Tx, model models.TransactionModel) bool {
	sqlTransaction := "INSERT INTO transaction(total_price, transaction_date, user_id, id, user_email, user_mobile, user_name) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7)"

	today := time.Now()
	idTransaction, _ := uuid.NewUUID()
	result, err := tx.ExecContext(context, sqlTransaction, model.Total, today.Format("2006-01-02 15:04:05"), model.UserId, idTransaction, "", "", model.Name)
	helpers.PanicIfError(err)

	var statement []string
	var values []interface{}
	for index, dt := range model.DetailTransaction {
		statement = append(statement, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			index*9+1,
			index*9+2,
			index*9+3,
			index*9+4,
			index*9+5,
			index*9+6,
			index*9+7,
			index*9+8,
			index*9+9))

		id, err := uuid.NewUUID()
		helpers.PanicIfError(err)
		values = append(values, idTransaction, dt.Id, id, dt.PerUnit, dt.Price, 1, dt.Weight, dt.ImageUrl, dt.Name)
	}

	sqlDetailTransaction := fmt.Sprintf("INSERT INTO detail_transaction(transaction_id,product_id,id,per_unit,price,total,weight,image_url,name) "+
		"VALUES %s", strings.Join(statement, ","))
	log.Default().Println(sqlDetailTransaction)

	result, err = tx.ExecContext(context, sqlDetailTransaction, values...)
	helpers.PanicIfError(err)

	//for _, dt := range model.DetailTransaction {
	//	id, err := uuid.NewUUID()
	//	helpers.PanicIfError(err)
	//
	//	result, err := tx.ExecContext(context, sqlDetailTransaction, idTransaction, dt.Id, id, dt.PerUnit, dt.Price, 1, dt.Weight, dt.ImageUrl, dt.Name)
	//	helpers.PanicIfError(err)
	//
	//	affected, err := result.RowsAffected()
	//	helpers.PanicIfError(err)
	//
	//	if affected > 0 {
	//		panic("err add detail transaction")
	//	}
	//}

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
