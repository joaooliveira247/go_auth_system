package errors

import (
	"errors"
	"fmt"
)

type BaseError struct {
	Resource string
	Msg      string
}

var (
	ErrNotExpected     = errors.New("NotExpectedTestError")
	ErrNothingToUpdate = errors.New("NothingToUpdate")
	ErrNothingToDelete = errors.New("NothingToDelete")
)

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

func NewModelError(err error) BaseError {
	return BaseError{
		Resource: "Model", Msg: err.Error(),
	}
}

func NewTokenSignUpError(err error) BaseError {
	return BaseError{
		Resource: "SignUp Token",
		Msg:      err.Error(),
	}
}
