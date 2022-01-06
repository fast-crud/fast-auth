package middleware

import (
	"encoding/base64"
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/auth"
	model "github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/app/service/basic"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"strings"
)

// Authentication
// @Description:  token认证
// @param r
//
func Authentication(r *ghttp.Request) {
	var needAuth = GetRequestAnnotationAuth(r)
	if needAuth == "false" {
		//无需auth
		r.Middleware.Next()
		return
	}
	tokenStr := r.Request.Header.Get(constants.HeaderAuth)
	if tokenStr == "" {
		err := gerror.NewCode(constants.CodeNoAuth, "还未登录")
		r.SetCtxVar("error", err)
		return
	}
	var tokenSplit = strings.Split(tokenStr, " ")
	var tokenType = tokenSplit[0]
	var token = tokenSplit[1]
	var claims *auth.Claims
	var err error
	if tokenType == "Bearer" {
		claims, err = resolveJwtToken(token)
	} else {
		//basicToken  app token， rpc调用
		claims, err = resolveBasicToken(token)
	}
	if err != nil {
		r.SetCtxVar("error", err)
		return
	}
	if claims == nil {
		err = gerror.NewCode(constants.CodeTokenResolveError, "token解析失败")
		r.SetCtxVar("error", err)
		return
	}
	r.SetCtxVar(constants.CtxAuth, claims)
	r.Middleware.Next()
}

//app授权
func resolveBasicToken(token string) (*auth.Claims, error) {
	clientId, clientSecret, ok := parseBasicAuth(token)
	if !ok || clientId == "" || clientSecret == "" {
		return nil, gerror.NewCode(constants.CodeTokenInvalid)
	}
	var key = constants.CacheApplicationByClientId + ":" + clientId
	value, _ := gcache.Get(gctx.New(), key)
	var app *model.Application
	if value != nil {
		*app = value.Val().(model.Application)
	}
	if app == nil {
		//缓存中没有，从数据库中获取
		app, _ = system.ApplicationService.GetByClientId(clientId)
	}

	//没有对应的应用
	if app == nil || app.Id == 0 {
		return nil, gerror.NewCode(constants.CodeTokenInvalid)
	}
	//secret错误
	if app.ClientSecret != clientSecret {
		return nil, gerror.NewCode(constants.CodeTokenInvalid)
	}

	var claims = auth.Claims{
		Username: app.ClientId,
		Type:     "basic",
		From:     "app",
	}
	return &claims, nil

}

func parseBasicAuth(token string) (username, password string, ok bool) {
	// Case insensitive prefix match. See Issue 22736.
	c, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

func resolveJwtToken(token string) (*auth.Claims, error) {
	_jwt := basic.NewJWT()
	claims, err := _jwt.ParseToken(token)
	if err != nil {
		if err == basic.TokenExpired {
			return nil, gerror.NewCode(constants.CodeTokenExpired, "token已过期")
		}
		err := gerror.NewCode(constants.CodeTokenResolveError, "token解析失败")
		return nil, err
	}
	return claims, nil
}
