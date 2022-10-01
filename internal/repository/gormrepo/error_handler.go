package gormrepo

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/pkg/mysqlerr"
	"gorm.io/gorm"
)

// Handler error for support gorm
func errHandler(err error) error {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
			return errs.ErrEntriyDuplicate
		}

		return errors.New("mysql driver error not found")
	}

	switch err.Error() {
	case gorm.ErrRecordNotFound.Error():
		return errs.ErrRecordNotFound
	default:
		return err
	}

}
