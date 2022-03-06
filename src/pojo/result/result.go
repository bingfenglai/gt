package result

import "log"

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

	return Result{
		Code:    1,
		Message: "fail",
		Data:    data,
	}
}

func FailWithErr(err error) Result {

	return Result{
		Code:    1,
		Message: "fail",
		Data:    err.Error(),
	}
}

func FailWithMsg(msg string, data interface{}) Result {

	return Result{
		Code:    1,
		Message: msg,
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
