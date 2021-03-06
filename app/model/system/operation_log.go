package system

import (
	"time"

	"github.com/fast-crud/fast-auth/library/global"
)

type OperationLog struct {
	global.Model
	Status       int           `json:"status" gorm:"column:status;comment:请求状态"`
	Ip           string        `json:"ip" gorm:"column:ip;comment:请求ip"`
	Path         string        `json:"path" gorm:"column:path;comment:请求路径"`
	Method       string        `json:"method" gorm:"column:method;comment:请求方法"`
	Agent        string        `json:"agent" gorm:"column:agent;comment:代理"`
	Latency      time.Duration `json:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"`
	Request      string        `json:"body" gorm:"type:text;column:request;comment:请求body"`
	Response     string        `json:"resp" gorm:"type:text;column:response;comment:响应Body"`
	ErrorMessage string        `json:"errorMessage" gorm:"column:error_message;comment:错误信息"`
	UserId       uint          `json:"userId" gorm:"column:user_id;comment:用户id"`
}

func (o *OperationLog) TableName() string {
	return "a_operation_log"
}
