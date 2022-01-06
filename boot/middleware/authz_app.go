package middleware

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/auth"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

//
// AuthzApp
// @Description: rpc鉴权
// @param r
//
func AuthzApp(r *ghttp.Request) {
	authClaims := r.GetCtxVar(constants.CtxAuth)
	if authClaims.Val() == nil {
		r.Middleware.Next()
		return
	}
	var loginUser auth.Claims
	if err := authClaims.Struct(&loginUser); err != nil {
		err := gerror.NewCode(constants.CodeNoPermission)
		r.SetCtxVar("error", err)
		return
	}
	if loginUser.Type != "app" {
		err := gerror.NewCode(constants.CodeNoPermission)
		r.SetCtxVar("error", err)
		return
	}

	r.Middleware.Next()
	return
}
