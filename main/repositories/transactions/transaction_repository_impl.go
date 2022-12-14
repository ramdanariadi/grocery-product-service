package transactions

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"log"
	"strings"
	"time"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func (repository TransactionRepositoryImpl) FindByTransactionId(context context.Context, tx *sql.Tx, id string) (*sql.Row, *sql.Rows) {
	queryTransaction := "SELECT id, total_price, transaction_date " +
		"FROM transaction " +
		"WHERE id = $1"
	rows := tx.QueryRowContext(context, queryTransaction, id)

	queryDetailTransaction := "SELECT name, id, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM detail_transaction " +
		"WHERE transaction_id = $1"
	dtRows, err := tx.QueryContext(context, queryDetailTransaction, id)
	helpers.PanicIfError(err)
	return rows, dtRows
}

func (repository TransactionRepositoryImpl) FindByUserId(context context.Context, tx *sql.Tx, userId string) (*sql.Rows, *sql.Rows) {
	sqlTransaction := "SELECT id, total_price, transaction_date " +
		"FROM transaction WHERE user_id = $1"
	rows, err := tx.QueryContext(context, sqlTransaction, userId)
	helpers.PanicIfError(err)

	sqlDetailTransaction := "SELECT name, dt.id, image_url, product_id, price, weight, per_unit, total, transaction_id " +
		"FROM detail_transaction dt " +
		"JOIN transaction t ON t.id = dt.transaction_id " +
		"WHERE t.user_id = $1"
	rowDetailTransaction, err := tx.QueryContext(context, sqlDetailTransaction, userId)
	helpers.PanicIfError(err)

	return rows, rowDetailTransaction
}

func (repository TransactionRepositoryImpl) Save(context context.Context, tx *sql.Tx, model models.TransactionModel) {
	sqlTransaction := "INSERT INTO transaction(id, total_price, transaction_date, user_id) VALUES($1,$2,$3,$4)"

	today := time.Now()
	idTransaction, _ := uuid.NewUUID()
	_, err := tx.ExecContext(context, sqlTransaction, idTransaction, model.TotalPrice, today.Format("2006-01-02 15:04:05"), model.UserId)
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

	_, err = tx.ExecContext(context, sqlDetailTransaction, values...)
	helpers.PanicIfError(err)
}

func (repository TransactionRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) {
	sqlDetailTransaction := "DELETE FROM detail_transaction WHERE id_transaction = $1"
	result, err := tx.ExecContext(context, sqlDetailTransaction, id)
	helpers.PanicIfError(err)

	_, err = result.RowsAffected()
	helpers.PanicIfError(err)

	sqlTransaction := "DELETE FROM transaction WHERE id = $1"
	execContext, err := tx.ExecContext(context, sqlTransaction, id)
	helpers.PanicIfError(err)

	_, err = execContext.RowsAffected()
	helpers.PanicIfError(err)
}
