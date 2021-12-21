package util

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/auth"
	"github.com/gogf/gf/v2/net/ghttp"
	"go.uber.org/zap"
)

var Auth = new(claims)

type claims struct{}

func (c *claims) GetAuthInfo(r *ghttp.Request) *auth.Claims {
	data := r.GetCtxVar(constants.CtxAuth)
	var _claims auth.Claims
	if err := data.Struct(&_claims); err != nil {
		zap.L().Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!", zap.Error(err))
	}
	return &_claims
}
