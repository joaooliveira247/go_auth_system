package errors

import "fmt"

type BaseError struct {
	Resource string
	Msg      string
}

func (err *BaseError) Error() string {
	return fmt.Sprintf("(%s): %s", err.Resource, err.Msg)
}
