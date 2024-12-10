package errors

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string) Error {
	return Error{Code: code, Message: message}
}

func NewErrorf(code int, format string, args ...interface{}) Error {
	return Error{Code: code, Message: fmt.Sprintf(format, args...)}
}

var (
	ErrInvalidCredentials = Error{Code: 401, Message: "Invalid credentials"}
)
