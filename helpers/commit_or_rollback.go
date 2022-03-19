package helpers

import (
	"database/sql"
	"reflect"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	if err, ok := err.(error); ok {
		panic(reflect.TypeOf(err))
	}

}
