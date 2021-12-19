package middleware

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"
	"reflect"
	"strings"
	"unsafe"
)

func GetHandlerAnnotation(r *ghttp.Request, annotationName string) string {
	var handles = reflect.ValueOf(r).Elem().FieldByName("handlers")
	var length = handles.Len()
	if length <= 0 {
		return ""
	}
	var lastHandler = handles.Index(length - 1)
	var Handler = lastHandler.Elem().FieldByName("Handler")
	var HandlerName = Handler.Elem().FieldByName("Name")
	var HandlerNameStr = reflect.NewAt(HandlerName.Type(), unsafe.Pointer(HandlerName.UnsafeAddr())).Elem()
	if strings.IndexAny(HandlerNameStr.String(), "controller") < 0 {
		return ""
	}
	var Info = Handler.Elem().FieldByName("Info")
	var Type = Info.FieldByName("Type")
	//获取私有属性
	var reqType = reflect.NewAt(Type.Type(), unsafe.Pointer(Type.UnsafeAddr())).Elem()
	var t = reqType.Interface().(reflect.Type)
	var objectReq = reflect.New(t.In(1))
	//拿到xxxReq的类定义，提取注解
	var v = gmeta.Get(objectReq, annotationName)
	var per = v.String()
	return per
}

func GetHandlerPermission(r *ghttp.Request) string {
	return GetHandlerAnnotation(r, constants.AnnoPermission)
}

func GetHandlerAuth(r *ghttp.Request) string {
	return GetHandlerAnnotation(r, constants.AnnoAuth)
}
