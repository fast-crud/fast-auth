package system

import (
	"context"
	"time"

	"github.com/fast-crud/fast-auth/library/global"
)

var JwtBlacklist = new(jwtBlacklist)

type jwtBlacklist struct{}

// JwtToBlacklist 拉黑jwt
// Author: [SliverHorn](https://github.com/SliverHorn)
//func (s *jwtBlacklist) JwtToBlacklist(jwt string) error {
//	entity := system.JwtBlacklist{Jwt: jwt}
//	if err := global.Db.Create(&entity).Error; err != nil {
//		return errors.Wrap(err, "拉黑jwt失败!")
//	}
//	global.JwtCache.SetDefault(jwt, struct{}{})
//	return nil
//}

// IsBlacklist 判断JWT是否在jwt黑名单
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *jwtBlacklist) IsBlacklist(jwt string) bool {
	_, ok := global.JwtCache.Get(jwt)
	return ok
}

// GetRedisJWT 获取用户在Redis的token
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *jwtBlacklist) GetRedisJWT(uuid string) (string, error) {
	return global.Redis.Get(context.Background(), uuid).Result()
}

// SetRedisJWT 保存jwt到Redis
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *jwtBlacklist) SetRedisJWT(uuid string, jwt string) error {
	timer := time.Duration(global.Config.Jwt.ExpiresTime) * time.Second
	return global.Redis.Set(context.Background(), uuid, jwt, timer).Err()
}

// ValidatorRedisToken 鉴权jwt
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *jwtBlacklist) ValidatorRedisToken(userUUID string, oldToken string) bool {
	if jwt, err := s.GetRedisJWT(userUUID); err != nil {
		return false
	} else {
		if jwt != oldToken {
			return false
		}
		return true
	}
}

// LoadJwt 加载jwt黑名单到 global.JwtCache 中
// Author [SliverHorn](https://github.com/SliverHorn)
//func (s *jwtBlacklist) LoadJwt() {
//	var data []string
//	err := global.Db.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
//	if err != nil {
//		zap.L().Error("加载失败!", zap.Error(err))
//	}
//	for i := range data { // 从db加载jwt数据
//		global.JwtCache.SetDefault(data[i], struct{}{})
//	}
//}
