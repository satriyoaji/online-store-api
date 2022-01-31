package helper

import (
	"github.com/jinzhu/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	if err := recover(); err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback.Error)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit.Error)
	}
}
