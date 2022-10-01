package errs

func HttpError(err error) error {
	switch err.Error() {
	case ErrRecordNotFound.Error():
		return NotFound(ErrRecordNotFound.Error())
	case ErrEntriyDuplicate.Error():
		return Conflict("this record has already exist")
	default:
		return Unexpected()
	}
}
