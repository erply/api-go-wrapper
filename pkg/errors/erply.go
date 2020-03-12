package errors

import "fmt"

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
