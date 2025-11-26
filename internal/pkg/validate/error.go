package validate

import "errors"

func NewError(errStr string) error {
	return errors.New(errStr)
}
