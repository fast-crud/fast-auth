package oss

import (
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/fast-crud/fast-auth/library/interfaces"
)

func Oss() interfaces.Oss {
	switch global.Config.System.OssType {
	case "local":
		return Local
	case "qiniu":
		return Qiniu
	case "minio":
		return Minio
	case "aliyun":
		return Aliyun
	case "tencent":
		return Tencent
	default:
		return Local
	}
}
