package result

import "log"

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

func NewResult(code int, msg string, data interface{}) Result {

	return Result{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func Ok(data interface{}) Result {

	return Result{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func OkWithMsg(data interface{}, msg string) Result {
	return Result{
		Code:    0,
		Message: msg,
		Data:    data,
	}
}

func Fail(data interface{}) Result {
	log.Default().Fatal()
	return Result{
		Code:    1,
		Message: "fail",
		Data:    data,
	}
}

func FailWithNilData() Result {
	log.Default().Fatal()
	return Result{
		Code:    1,
		Message: "fail",
		Data:    nil,
	}
}
