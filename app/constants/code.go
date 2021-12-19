package constants

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeNoAuth            = gcode.New(1001, "您还未登录", nil)
	CodeNoPermission      = gcode.New(1002, "权限不足！", nil)
	CodeTokenExpired      = gcode.New(1003, "您的Token已过期，请重新登录", nil)
	CodeTokenResolveError = gcode.New(1004, "Token解析失败", nil)
	CodeTokenInvalid      = gcode.New(1005, "您的Token已失效", nil)
	CodeUserNotExists     = gcode.New(1006, "用户不存在", nil)
	CodeGetTokenFailed    = gcode.New(1007, "获取Token失败", nil)
	CodeParseTokenFailed  = gcode.New(1008, "解析Token失败", nil)
)

var (
	CodeParamCantBlank = gcode.New(2001, "参数不能为空", nil)
	CodeUserFindError  = gcode.New(2002, "查询用户失败", nil)
)
