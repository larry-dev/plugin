package response

import "gitee.com/egcode/plugins/ginx/exception"

//2	                        05	           02
//服务级错误（1为系统级错误）	服务模块代码	   具体错误代码
var (
	// success code
	OK = &exception.BaseException{Code: 0, Message: "OK"}
	// system error codes
	InternalServerError = &exception.BaseException{Code: 10001, Message: "Internal server error."}
	ErrBind             = &exception.BaseException{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// service error codes
	ErrUserNotFound = &exception.BaseException{Code: 20102, Message: "The user was not found."}
)

/**
 * 格式化异常数据
 */
func DecodeException(err interface{}) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *exception.Exception:
		return typed.Code, typed.Message
	case *exception.BaseException:
		return typed.Code, typed.Message
	case error:
		return InternalServerError.Code, typed.Error()
	}
	return InternalServerError.Code, InternalServerError.Message
}
