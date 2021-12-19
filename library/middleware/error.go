package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var ctx = r.GetCtx()
	var (
		err error
		res interface{}
	)
	res, err = r.GetHandlerResponse()
	if err != nil {
		g.Log().Error(ctx, "请求出错!", g.Map{"err": err})
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		writeResponse(r, &ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: err.Error(),
			Data:    nil,
		})
		return
	} else {
		var err1 = r.GetCtxVar("error").Val()
		if err1 != nil {
			var err = err1.(error)
			g.Log().Error(ctx, "请求出错!", g.Map{"err": err})
			code := gerror.Code(err)
			if code == gcode.CodeNil {
				code = gcode.CodeInternalError
			}
			writeResponse(r, &ghttp.DefaultHandlerResponse{
				Code:    code.Code(),
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
	}
	writeResponse(r, &ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: "",
		Data:    res,
	})

}

func writeResponse(r *ghttp.Request, res *ghttp.DefaultHandlerResponse) {
	var internalErr = r.Response.WriteJson(res)

	if internalErr != nil {
		g.Log().Error(r.GetCtx(), "response error!", g.Map{"err": internalErr})
	}
}
