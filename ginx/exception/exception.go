package exception

import "fmt"

type BaseException struct {
	Code    int
	Message string
}

func (e *BaseException) Error() string {
	return e.Message
}
func (e *BaseException) RMessage(msg string) *BaseException {
	e.Message = msg
	return e
}

type Exception struct {
	Code    int
	Message string
	Err     error
}

func New(ex *BaseException, err error) *Exception {
	return &Exception{ex.Code, ex.Message, err}
}

func (ex *Exception) Add(message string) error {
	//err.Message = fmt.Sprintf("%s %s", err.Message, message)
	ex.Message += " " + message
	return ex
}

func (ex *Exception) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	ex.Message += " " + fmt.Sprintf(format, args...)
	return ex
}

func (ex *Exception) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", ex.Code, ex.Message, ex.Err)
}
