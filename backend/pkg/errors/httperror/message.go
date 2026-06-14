package httperror

type ErrorMessage string

func (e ErrorMessage) String() string { return string(e) }

const (
	MsgBadRequest     ErrorMessage = "bad request"
	MsgNotFound       ErrorMessage = "resource not found"
	MsgAlreadyExists  ErrorMessage = "resource already exists"
	MsgInternalError  ErrorMessage = "internal server error"
	MsgNotImplemented ErrorMessage = "not implemented"
)
