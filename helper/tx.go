package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
		if err != nil {
			errorRollbacck := tx.Rollback()
			PanicIfError(errorRollbacck)
			panic(err)
		}else{
			errorCode := tx.Commit()
			PanicIfError(errorCode)
		}
}