package system

import (
	"github.com/fast-crud/fast-auth/app/constants"
	auth2 "github.com/fast-crud/fast-auth/app/model/auth"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	"github.com/fast-crud/fast-auth/app/service/basic"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"

	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/common"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var UserService = new(userService)

type userService struct{}

type UserRegisterParams struct {
	Avatar   string
	Username string
	Password string
	NickName string
}

// Register
// @Description: 用户注册
// @receiver userService
// @param info
// @return data
// @return err
func (userService *userService) Register(info *UserRegisterParams) (data *system.User, err error) {
	var entity system.User
	if !errors.Is(global.Db.Where("username = ?", info.Username).First(&entity).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return nil, errors.Wrap(err, "用户名已注册")
	}
	entity = system.User{
		Avatar:   info.Avatar,
		Username: info.Username,
		Password: info.Password,
		NickName: info.NickName,
	}
	if err = entity.EncryptedPassword(); err != nil {
		return nil, errors.Wrap(err, "密码加密失败!")
	}
	if err = global.Db.Create(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "用户注册失败!")
	}
	return &entity, nil
}

// Login
// @Description: 用户登录接口
// @receiver userService
// @param username
// @param password
// @return token
// @return err
//
func (userService *userService) Login(username string, password string) (token *res.AccessTokenRes, err error) {
	var entity system.User
	if errors.Is(global.Db.Where("username = ?", username).Preload("Roles").First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, gerror.NewCode(constants.CodeUserNotExists)
	}
	if !entity.CompareHashAndPassword(password) {
		return nil, errors.New("密码错误!")
	}
	return userService.tokenCreate(&entity)
}

//
// GetById
// @Description: 根据id获取用户信息
// @receiver userService
// @param Id
// @return data
// @return err
//
func (userService *userService) GetById(Id uint) (data *system.User, err error) {
	if Id == 0 {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	var entity system.User
	var search = func(db *gorm.DB) *gorm.DB {
		if Id != 0 {
			db = db.Where("id = ?", Id)
		}
		return db
	}
	if err = global.Db.Scopes(search).Preload("Roles").First(&entity).Error; err != nil {
		return nil, gerror.NewCode(constants.CodeUserFindError)
	}
	return &entity, nil
}

type UserFindParams struct {
	Id       uint   `json:"id" example:"7"`
	Username string `json:"username" example:"7"`
}

//
//  Find
//  @Description: 查询user
//  @receiver userService
//  @param userFindParams
//  @return data
//  @return err
//
func (userService *userService) Find(userFindParams *UserFindParams) (data *system.User, err error) {
	var entity system.User

	if userFindParams.Username == "" && userFindParams.Id == 0 {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	var search = func(db *gorm.DB) *gorm.DB {
		if userFindParams.Id != 0 {
			db = db.Where("id = ?", userFindParams.Id)
		}
		if userFindParams.Username != "" {
			db = db.Where("username = ?", userFindParams.Username)
		}
		return db
	}
	if err = global.Db.Scopes(search).Preload("Roles").First(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "用户查询失败!")
	}
	return &entity, nil
}

type UserUpdateParams struct {
	Id       uint   `json:"id" example:"uint 主键ID"`
	Avatar   string `json:"avatar" example:"用户头像"`
	NickName string `json:"nickName" example:"用户昵称"`
}

// Update
// @Description: 更新用户信息
// @receiver userService
// @param userUpdateParams
// @return user
// @return err
//
func (userService *userService) Update(userUpdateParams *UserUpdateParams) (user *system.User, err error) {
	update := system.User{
		Avatar:   userUpdateParams.Avatar,
		NickName: userUpdateParams.NickName,
	}
	err = global.Db.Where("id = ?", userUpdateParams.Id).Updates(&update).Error
	return &update, err
}

type UserChangePasswordParams struct {
	Id          uint   `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// ChangePassword
// @Description: 修改密码
// @receiver userService
// @param params
// @return error
func (userService *userService) ChangePassword(params *UserChangePasswordParams) error {
	var entity system.User
	err := global.Db.Where("id = ?", params.Id).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "用户不存在! ")
	}
	if !entity.CompareHashAndPassword(params.Password) {
		return errors.Wrap(err, "密码错误!")
	}
	entity.Password = params.NewPassword
	if err = entity.EncryptedPassword(); err != nil {
		return errors.Wrap(err, "密码加密失败!")
	}
	return global.Db.Where("id = ?", params.Id).Update("password", entity.Password).Error
}

// RestPassword
// @Description: 重置密码
// @receiver userService
// @param params
// @return error
func (userService *userService) ResetPassword(Id uint, NewPassword string) error {
	var entity system.User
	err := global.Db.Where("id = ?", Id).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "用户不存在! ")
	}
	entity.Password = NewPassword
	if err = entity.EncryptedPassword(); err != nil {
		return errors.Wrap(err, "密码加密失败!")
	}
	return global.Db.Model(&entity).Where("id = ?", Id).Update("password", entity.Password).Error
}

func (userService *userService) Delete(Id uint) error {
	return global.Db.Delete(&system.User{}, Id).Error
}

// GetList 获取用户列表
func (userService *userService) GetList(info *common.PageInfo) (list []system.User, total int64, err error) {
	entities := make([]system.User, 0, info.PageSize)
	db := global.Db.Model(&system.User{})
	err = db.Count(&total).Error
	err = db.Scopes(common.Paginate(info)).Preload("Authority").Preload("Authorities").Find(&entities).Error
	return entities, total, err
}

//
//  tokenCreate
//  @Description: token生成
//  @receiver s
//  @param user
//  @return *response.UserLogin
//  @return error
//
func (userService *userService) tokenCreate(user *system.User) (*res.AccessTokenRes, error) {
	_jwt := basic.NewJWT()
	var roleIds = make([]uint, len(user.Roles))
	for i := 0; i < len(user.Roles); i++ {
		roleIds[i] = user.Roles[i].Id
	}
	claims := auth2.Claims{
		Id:         user.Id,
		Username:   user.Username,
		RoleIds:    roleIds,
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
