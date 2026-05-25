// Package httperror provides HTTP error response types and constants.
package httperror

// ErrorCode represents an HTTP error code identifier.
type ErrorCode string

func (e ErrorCode) String() string { return string(e) }

const (
	CodeBadRequest     ErrorCode = "BAD_REQUEST"
	CodeNotFound       ErrorCode = "NOT_FOUND"
	CodeAlreadyExists  ErrorCode = "ALREADY_EXISTS"
	CodeInternalError  ErrorCode = "INTERNAL_ERROR"
	CodeNotImplemented ErrorCode = "NOT_IMPLEMENTED"
)
