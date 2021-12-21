package middleware

import (
	"bytes"
	"github.com/fast-crud/fast-auth/app/controller/util"
	"io"
	"time"

	model "github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/net/ghttp"
	"go.uber.org/zap"
)

func OperationLog(r *ghttp.Request) {
	// Request
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		zap.L().Error("读取内容失败", zap.Error(err))
	}

	r.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	var auth = util.Auth.GetAuthInfo(r)

	record := system.OperationLogCreateParams{
		OperationLog: model.OperationLog{
			Ip:      r.GetClientIp(),
			Method:  r.Request.Method,
			Path:    r.Request.URL.Path,
			Agent:   r.Request.UserAgent(),
			Request: string(body),
			UserId:  auth.Id,
		},
	}
	now := time.Now()

	r.Middleware.Next()

	// Response
	latency := time.Now().Sub(now)

	if err = r.GetError(); err != nil {
		record.ErrorMessage = err.Error()
	}

	record.Status = r.Response.Status
	record.Latency = time.Duration(latency.Microseconds())
	//record.Response = string(r.Response.Buffer())

	if err = system.OperationLog.Create(&record); err != nil {
		zap.L().Error("创建日志记录失败!", zap.Error(err))
	}
}
