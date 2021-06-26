package customeError

type CustomeError struct {
	ErrorCode    int
	ErrorMessage string
}

func New(code int, message string) error {
	return &CustomeError{ErrorCode: code, ErrorMessage: message}
}

func (e *CustomeError) Error() string {
	return e.ErrorMessage
}
