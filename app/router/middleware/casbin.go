package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/protocol/goai"
	"github.com/gogf/gf/v2/util/gmeta"
	"reflect"
)

// Casbin Casbin拦截器
func Casbin(r *ghttp.Request) {
	//claims := r.GetCtxVar("claims")
	//var user request.CustomClaims
	//if err := claims.Struct(&user); err != nil {
	//	_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "权限不足!"})
	//	r.ExitAll()
	//	return
	//}
	var handlerReflect = reflect.ValueOf(r).Elem().FieldByName("handlers")
	var last = handlerReflect.Len()
	var lastHandler = handlerReflect.Index(last - 1)
	var Handler = lastHandler.Elem().FieldByName("Handler")
	fmt.Printf("type-kind: %s, Name:%s\n", Handler.Elem().Kind(), Handler.Elem().FieldByName("Name"))
	var info = Handler.Elem().FieldByName("Info")
	var Type = info.FieldByName("Type").Interface().(reflect.Type)
	var (
		objectReq = reflect.New(Type.In(1))
	)
	var v = gmeta.Get(objectReq, goai.TagNamePath)
	fmt.Printf("type-kind: %s, %s\n", v)
	r.Middleware.Next()
	//method := r.Request.Method        // 获取请求方法
	//authorityId := user.AuthorityId   // 获取用户的角色
	//e := system.Casbin.Casbin()
	//success, _ := e.Enforce(authorityId, url, method)
	//if global.Config.System.Env == "develop" || success {
	//	r.Middleware.Next()
	//} else {
	//	_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "权限不足!"})
	//	r.ExitAll()
	//	return
	//} // 判断策略中是否存在
	return
}
