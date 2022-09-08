package errs

import (
	"errors"
	"fmt"
)

func NaN(field ...string) error {
	if len(field) > 0 {
		return fmt.Errorf("%s is not a number", field[0])
	}

	return errors.New("not a number")
}
