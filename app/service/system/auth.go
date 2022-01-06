package system

import (
	"github.com/fast-crud/fast-auth/app/constants"
	auth2 "github.com/fast-crud/fast-auth/app/model/auth"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	"github.com/fast-crud/fast-auth/app/service/basic"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"

	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var AuthService = new(userService)

type authService struct{}

// Login
// @Description: 用户登录接口
// @receiver userService
// @param username
// @param password
// @return token
// @return err
//
func (authService *authService) Login(username string, password string) (token *res.AccessTokenRes, err error) {
	var entity system.User
	if errors.Is(global.Db.Where("username = ?", username).Preload("Roles").First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, gerror.NewCode(constants.CodeUserNotExists)
	}
	if !entity.CompareHashAndPassword(password) {
		return nil, errors.New("密码错误!")
	}
	return authService.createUserToken(&entity)
}

// appToken
// @Description: appLogin接口
// @receiver authService
// @param username
// @param password
// @return token
// @return err
//
func (authService *authService) appToken(clientId string, clientSecret string) (token *res.AccessTokenRes, err error) {
	var entity system.Application
	if errors.Is(global.Db.Where("clientId = ?", clientId).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, gerror.NewCode(constants.CodeAppNotExists)
	}
	if !entity.CompareHashAndSecret(clientSecret) {
		return nil, errors.New("密码错误!")
	}
	return authService.createAppToken(&entity)
}

//
//  createAppToken
//  @Description: token生成
//  @receiver s
//  @param user
//  @return error
//
func (authService *authService) createUserToken(user *system.User) (*res.AccessTokenRes, error) {
	_jwt := basic.NewJWT()
	var roleIds = make([]uint, len(user.Roles))
	for i := 0; i < len(user.Roles); i++ {
		roleIds[i] = user.Roles[i].Id
	}
	claims := auth2.Claims{
		Id:         user.Id,
		Username:   user.Username,
		RoleIds:    roleIds,
		Type:       "user",
		BufferTime: global.Config.Jwt.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Config.Jwt.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "handsfree",                                       // 签名的发行者
		},
	}
	token, err := _jwt.CreateToken(&claims)
	if err != nil {
		return nil, errors.Wrap(err, "获取token失败!")
	}
	entity := res.AccessTokenRes{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
	return &entity, nil
}

//
//  createAppToken
//  @Description: token生成
//  @receiver s
//  @param user
//  @return *response.UserLogin
//  @return error
//
func (authService *authService) createAppToken(app *system.Application) (*res.AccessTokenRes, error) {
	_jwt := basic.NewJWT()
	claims := auth2.Claims{
		Id:         app.Id,
		Username:   app.ClientId,
		Type:       "app",
		BufferTime: global.Config.Jwt.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Config.Jwt.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "handsfree",                                       // 签名的发行者
		},
	}
	token, err := _jwt.CreateToken(&claims)
	if err != nil {
		return nil, errors.Wrap(err, "获取token失败!")
	}
	entity := res.AccessTokenRes{App: app, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
	return &entity, nil
}
