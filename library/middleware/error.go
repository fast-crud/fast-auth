package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"go.uber.org/zap"
)

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var err1 = r.GetCtxVar("error").Val()
	if err1 != nil {
		var err = err1.(error)
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		var internalErr = r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: err.Error(),
			Data:    nil,
		})

		if internalErr != nil {
			zap.L().Error("异常处理失败", zap.Error(internalErr))
		}
		return
	}
}
