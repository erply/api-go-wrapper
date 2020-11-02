package errors

import (
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api/common"
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

func NewErplyError(status, msg string) *ErplyError {
	return &ErplyError{Status: status, Message: msg}
}

func NewErplyErrorf(status, msg string, args ...interface{}) *ErplyError {
	return &ErplyError{Status: status, Message: fmt.Sprintf(msg, args...)}
}

func NewFromResponseStatus(status *common.Status) *ErplyError {
	var s string
	if status.ErrorField != "" {
		s = fmt.Sprintf("%s, error field: %s", status.ErrorCode.String(), status.ErrorField)
	} else {
		s = status.ErrorCode.String()
	}
	m := status.Request + ": " + status.ResponseStatus
	return &ErplyError{Status: s, Message: m}
}

func NewFromError(msg string, err error) *ErplyError {
	if err != nil {
		return NewErplyError("Error", errors.Wrap(err, msg).Error())
	}
	return NewErplyError("Error", msg)
}
