package system

import "github.com/fast-crud/fast-auth/library/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (j *JwtBlacklist) TableName() string {
	return "a_jwt_blacklist"
}
