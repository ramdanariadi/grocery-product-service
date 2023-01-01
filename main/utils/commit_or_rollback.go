package utils

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	if err, ok := err.(error); ok {
		panic(err)
	}

}
