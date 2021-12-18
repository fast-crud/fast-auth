package middleware

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/auth"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Casbin Casbin拦截器
func Casbin(r *ghttp.Request) {
	authClaims := r.GetCtxVar(constants.CtxAuth)
	var loginUser auth.JwtClaims
	if err := authClaims.Struct(&loginUser); err != nil {
		err := gerror.NewCode(constants.CodeNoPermission)
		r.SetCtxVar("error", err)
		return
	}
	var permission = GetHandlerPermission(r)
	roleIds := loginUser.RoleIds
	e := system.Casbin.Casbin()
	var ok = false
	for roleId := range roleIds {
		success, _ := e.Enforce(roleId, permission)
		if success {
			ok = true
			break
		}
	}
	if ok {
		r.Middleware.Next()
		return
	}
	err := gerror.NewCode(constants.CodeNoPermission)
	r.SetCtxVar(constants.CtxError, err)
}
