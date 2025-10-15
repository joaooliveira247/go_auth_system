package errors

import (
	"errors"
	"fmt"
)

type BaseError struct {
	Resource string
	Msg      string
}

var ErrNotExpected = errors.New("NotExpectedTestError")

func (err BaseError) Error() string {
	return fmt.Sprintf("(%s): %s", err.Resource, err.Msg)
}

func NewDatabaseError(err error) BaseError {
	return BaseError{
		Resource: "Database",
		Msg:      err.Error(),
	}
}

func NewHashError(err error) BaseError {
	return BaseError{
		Resource: "Hash",
		Msg:      err.Error(),
	}
}
