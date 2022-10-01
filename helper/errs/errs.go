package errs

import "errors"

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrEntriyDuplicate = errors.New("this record has alreaday exist")
)
