package middleware

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/service/basic"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

/**
给rpc的basic鉴权
*/
func BasicAuth(r *ghttp.Request) {
	var needAuth = GetRequestAnnotationAuth(r)
	if needAuth == "false" {
		//无需auth
		r.Middleware.Next()
		return
	}
	token := r.Request.Header.Get(constants.HeaderAuth)
	if token == "" {
		err := gerror.NewCode(constants.CodeNoAuth, "还未登录")
		r.SetCtxVar("error", err)
		return
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	//if system.JwtBlacklist.IsBlacklist(token) {
	//	err := gerror.NewCode(constants.CodeTokenInvalid, "令牌失效")
	//	r.SetCtxVar("error", err)
	//	return
	//}
	_jwt := basic.NewJWT()
	claims, err := _jwt.ParseToken(token)
	if err != nil {
		if err == basic.TokenExpired {
			err := gerror.NewCode(constants.CodeTokenExpired, "token已过期")
			r.SetCtxVar("error", err)
			return
		}
		err := gerror.NewCode(constants.CodeTokenResolveError, "token解析失败")
		r.SetCtxVar("error", err)
		return
	}
	r.SetCtxVar(constants.CtxAuth, claims)
	r.Middleware.Next()
}
