package cerror

type CustomError struct {
	Code    int
	Message string
	Err     error
}

type ICustomError interface {
	Error() string
	WithErr(err error) ICustomError
	WithMsg(msg string) ICustomError
	WithCode(code int) ICustomError
}

func New(code int, message string, err error) ICustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) WithErr(err error) ICustomError {
	e.Err = err
	return e
}

func (e *CustomError) WithMsg(msg string) ICustomError {
	e.Message = msg
	return e
}

func (e *CustomError) WithCode(code int) ICustomError {
	e.Code = code
	return e
}
