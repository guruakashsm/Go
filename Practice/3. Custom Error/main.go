package main

/*
	This program demonstrates custom error handling in Go using a structured `Error` type.
	It showcases how to define, use, and return meaningful errors with additional context.
	The workflow is as follows:
		1. The `Error` type is defined with fields for an error code and a message, implementing the `error` interface.
		2. A global variable `ErrInvalidCredentials` is created as a reusable instance of the `Error` type.
		3. The `main` function calls `ReturnError`, which returns a predefined error.
		4. If an error is returned, the program prints the error message by calling the `Error()` method.
	This approach allows for structured, consistent, and descriptive error handling in Go programs.
*/

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrInvalidCredentials = &Error{Code: 400, Message: "Invalid credentials"}
)

func main() {
	err := ReturnError()
	if err != nil {
		println(err.Error())
	}
}

func ReturnError() error {
	return ErrInvalidCredentials
}
