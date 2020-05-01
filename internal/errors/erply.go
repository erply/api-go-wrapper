package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErplyError struct {
	error
	Status  string
	Message string
}

func (e *ErplyError) Error() string {
	return fmt.Sprintf("ERPLY API: %s status: %s", e.Message, e.Status)
}

func NewErplyError(status string, msg string) *ErplyError {
	return &ErplyError{Status: status, Message: msg}
}

func NewFromError(msg string, err error) *ErplyError {
	if err != nil {
		return NewErplyError("Error", errors.Wrap(err, msg).Error())
	}
	return NewErplyError("Error", msg)
}
