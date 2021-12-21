package middleware

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"
	"reflect"
	"strings"
	"unsafe"
)

func GetRequestAnnotation(r *ghttp.Request, annotationName string) string {
	HandlerFuncType := getHandlerInfo(r)
	if HandlerFuncType == nil {
		return ""
	}

	if HandlerFuncType.NumIn() < 2 {
		return ""
	}
	var objectReq = reflect.New(HandlerFuncType.In(1))
	//拿到xxxReq的类定义，提取注解
	var v = gmeta.Get(objectReq, annotationName)
	if v == nil {
		return ""
	}
	var per = v.String()
	return per
}

func getPrivateValue(Type reflect.Value) reflect.Value {
	return reflect.NewAt(Type.Type(), unsafe.Pointer(Type.UnsafeAddr())).Elem()
}

func GetResponseAnnotation(r *ghttp.Request, annotationName string) string {
	handlerFuncType := getHandlerInfo(r)
	if handlerFuncType == nil {
		return ""
	}
	var objectRes = reflect.New(handlerFuncType.Out(0))
	//拿到xxxReq的类定义，提取注解
	var v = gmeta.Get(objectRes, annotationName)
	if v == nil {
		return ""
	}
	var per = v.String()
	return per
}

func getHandlerInfo(r *ghttp.Request) reflect.Type {
	var handles = reflect.ValueOf(r).Elem().FieldByName("handlers")
	var length = handles.Len()
	if length <= 0 {
		return nil
	}
	var lastHandler = handles.Index(length - 1)
	var Handler = lastHandler.Elem().FieldByName("Handler")
	var HandlerName = Handler.Elem().FieldByName("Name")
	var HandlerNameStr = reflect.NewAt(HandlerName.Type(), unsafe.Pointer(HandlerName.UnsafeAddr())).Elem()
	if strings.IndexAny(HandlerNameStr.String(), "controller") < 0 {
		return nil
	}
	var Info = Handler.Elem().FieldByName("Info")
	var Type = Info.FieldByName("Type")
	//获取私有属性
	handlerFunc := getPrivateValue(Type)
	var handlerFuncType = handlerFunc.Interface().(reflect.Type)
	return handlerFuncType
}

func GetRequestAnnotationPermission(r *ghttp.Request) string {
	return GetRequestAnnotation(r, constants.AnnoPermission)
}

func GetRequestAnnotationAuth(r *ghttp.Request) string {
	return GetRequestAnnotation(r, constants.AnnoAuth)
}
