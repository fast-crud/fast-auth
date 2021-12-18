package middleware

import (
	"bytes"
	"github.com/fast-crud/fast-auth/app/constants"
	"io"
	"strconv"
	"time"

	model "github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/net/ghttp"
	"go.uber.org/zap"
)

func OperationRecord(r *ghttp.Request) {
	// Request
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		zap.L().Error("读取内容失败", zap.Error(err))
	}

	r.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	id, _ := strconv.Atoi(r.Request.Header.Get(constants.HeaderUserId))

	record := system.OperationRecordCreateParams{
		OperationRecord: model.OperationRecord{
			Ip:      r.GetClientIp(),
			Method:  r.Request.Method,
			Path:    r.Request.URL.Path,
			Agent:   r.Request.UserAgent(),
			Request: string(body),
			UserId:  id,
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
	record.Response = string(r.Response.Buffer())

	if err = system.OperationRecord.Create(&record); err != nil {
		zap.L().Error("创建日志记录失败!", zap.Error(err))
	}
}
