package errors

type Error struct {
	Err     error
	Code    int
	Message string
}

func (e Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	} else {
		return e.Message
	}
}

func NewError(code int, message string, err error) *Error {
	return &Error{Code: code, Message: message, Err: err}
}
