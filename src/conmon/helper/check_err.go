package helper

import "go.uber.org/zap"

func CheckErr(err error) (bool, string) {

	if err == nil {
		return true, ""
	}

	zap.L().Error(err.Error())
	return false, err.Error()
}
