package config

type Config struct {
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`             // jwt
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`             // zap 日志
	Gorm    Gorm    `mapstructure:"gorm" json:"gorm" yaml:"gorm"`          // gorm
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`       // redis
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`       // email 邮箱
	System  System  `mapstructure:"system" json:"system" yaml:"system"`    // system 系统
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`    // casbin 权限
	Casdoor Casdoor `mapstructure:"casdoor" json:"casdoor" yaml:"casdoor"` // casdoor 认证
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"` // captcha 验证码

	Qiniu     Qiniu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`                 // 七牛云对象存储
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`                 // 本地
	Minio     Minio     `mapstructure:"minio" json:"minio" yaml:"minio"`                 // minio对象存储
	Aliyun    Aliyun    `mapstructure:"aliyun" json:"aliyun" yaml:"aliyun"`              // 阿里云对象存储
	Tencent   Tencent   `mapstructure:"tencent" json:"tencent" yaml:"tencent"`           // 腾讯对象存储
	HuaWeiObs HuaWeiObs `mapstructure:"hua-wei-obs" json:"huaWeiObs" yaml:"hua-wei-obs"` // 华为云对象存储
}
