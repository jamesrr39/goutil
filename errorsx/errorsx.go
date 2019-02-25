package errorsx

import (
	"fmt"
	"runtime/debug"
	"strings"
)

type Error struct {
	OriginalError       error
	ExtraContextMessage string
	Stack               []byte
}

const indent = "\t"

func (e *Error) message() string {
	errMessage := e.OriginalError.Error()

	if e.ExtraContextMessage != "" {
		errMessage += "\n" + e.ExtraContextMessage
	}

	return errMessage
}

func (e *Error) Error() string {
	return e.message() + "\n" + e.prettyStackTrace()
}

func (e *Error) prettyStackTrace() string {
	stackLines := strings.Split(string(e.Stack), "\n")
	stackString := strings.Join(stackLines, "\n"+indent)

	return stackString
}

func newError(err error, extraContextMessage string) *Error {
	return &Error{
		err,
		extraContextMessage,
		debug.Stack(),
	}
}

func New(message string, args ...interface{}) *Error {
	return newError(fmt.Errorf(message, args...), "")
}

func Wrap(err error) *Error {
	return newError(err, "")
}

func Wrapf(err error, extraContextMessage string, args ...interface{}) *Error {
	return newError(err, fmt.Sprintf(extraContextMessage, args...))
}

func IsErrorx(err error) (*Error, bool) {
	errx, ok := err.(*Error)
	return errx, ok
}
