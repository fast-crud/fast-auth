package res

import "go.uber.org/zap"

type CommonRes struct {
	Code    int16
	Message string
	Data    interface{}
}

func Error(Message string, Error error) CommonRes {
	zap.L().Error(Message, zap.Error(Error))
	return CommonRes{Code: 1, Message: Message}
}
func ErrorWithCode(Message string, Code int16, Error error) CommonRes {
	zap.L().Error(Message, zap.Error(Error))
	return CommonRes{Code: Code, Message: Message}
}
func Success(Data interface{}) CommonRes {
	return CommonRes{Code: 0, Data: Data}
}
