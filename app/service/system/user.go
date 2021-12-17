package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/basic/res"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/auth"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var UserService = new(userService)

type userService struct{}

// Register
// @Description: 用户注册
// @receiver userService
// @param info
// @return data
// @return err
func (userService *userService) Register(info *request.UserRegister) (data *system.User, err error) {
	var entity system.User
	if !errors.Is(global.Db.Where("username = ?", info.Username).First(&entity).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return nil, errors.Wrap(err, "用户名已注册")
	}
	entity = info.Create()
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
func (userService *userService) Login(username string, password string) (token *res.AccessToken, err error) {
	var entity system.User
	if errors.Is(global.Db.Where("username = ?", username).Preload("Roles").First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在!")
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
		return nil, errors.New("id不能为空")
	}
	var entity system.User
	var search = func(db *gorm.DB) *gorm.DB {
		if Id != 0 {
			db = db.Where("id = ?", Id)
		}
		return db
	}
	if err = global.Db.Scopes(search).Preload("Roles").First(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "用户查询失败!")
	}
	return &entity, nil
}

type UserFindParams struct {
	Id       uint   `json:"id" example:"7"`
	username string `json:"username" example:"7"`
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

	if userFindParams.username == "" && userFindParams.Id == 0 {
		return nil, errors.New("查询条件不能为空")
	}
	var search = func(db *gorm.DB) *gorm.DB {
		if userFindParams.Id != 0 {
			db = db.Where("id = ?", userFindParams.Id)
		}
		if userFindParams.username != "" {
			db = db.Where("username = ?", userFindParams.username)
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

//// SetAuthority 设置用户的活跃角色
//// Author: [SliverHorn](https://github.com/SliverHorn)
//func (s *userService) SetAuthority(info *request.UserSetAuthority) error {
//	err := global.Db.Where("user_id = ? AND authority_id = ?", info.Id, info.AuthorityId).First(&system.UseAuthority{}).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return errors.New("该用户无此角色!")
//	}
//	if err = global.Db.Model(&system.User{}).Where("uuid = ?", info.Uuid).Update("authority_id", info.AuthorityId).Error; err != nil {
//		return errors.Wrap(err, "更新用户角色失败!")
//	}
//	return nil
//}

// SetUserAuthorities 设置用户可切换的角色
// Author [SliverHorn](https://github.com/SliverHorn)
//func (s *userService) SetUserAuthorities(info *request.UserSetAuthorities) error {
//	return global.Db.Transaction(func(tx *gorm.DB) error {
//		if err := tx.Delete(&[]system.UseAuthority{}, "user_id = ?", info.Id).Error; err != nil {
//			return errors.Wrap(err, "用户可切换的旧角色删除失败!")
//		}
//		length := len(info.AuthorityIds)
//		entities := make([]system.UseAuthority, 0, length)
//		for i := 0; i < length; i++ {
//			entities = append(entities, system.UseAuthority{UserId: info.Id, AuthorityId: info.AuthorityIds[i]})
//		}
//		if err := tx.Create(&entities).Error; err != nil {
//			return errors.Wrap(err, "设置用户多角色失败!")
//		}
//		return nil
//	})
//}

func (userService *userService) Delete(Id uint) error {
	return global.Db.Delete(&system.User{}, Id).Error
}

// GetList 获取用户列表
// Author: [SliverHorn](https://github.com/SliverHorn)
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
func (userService *userService) tokenCreate(user *system.User) (*res.AccessToken, error) {
	_jwt := auth.NewJWT()
	claims := request.CustomClaims{
		Id:       user.Id,
		NickName: user.NickName,
		Username: user.Username,
		//RoleIds: user.Roles.
		//AuthorityId: user.AuthorityId,
		BufferTime: global.Config.Jwt.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Config.Jwt.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                          // 签名的发行者
		},
	}
	token, err := _jwt.CreateToken(&claims)
	if err != nil {
		return nil, errors.Wrap(err, "获取token失败!")
	}
	if !global.Config.System.UseMultipoint {
		entity := res.AccessToken{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
		return &entity, nil
	}

	if jwtStr, _err := JwtBlacklist.GetRedisJWT(user.Username); _err == redis.Nil {
		if err = JwtBlacklist.SetRedisJWT(token, user.Username); err != nil {
			return nil, errors.Wrap(err, "设置登录状态失败!")
		}
		entity := res.AccessToken{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
		return &entity, nil
	} else if _err != nil {
		return nil, errors.Wrap(_err, "设置登录状态失败!")
	} else {
		if !JwtBlacklist.IsBlacklist(jwtStr) {
			return nil, errors.Wrap(_err, "jwt作废失败!")
		}
		if err = JwtBlacklist.SetRedisJWT(token, user.Username); err != nil {
			return nil, errors.Wrap(err, "设置登录状态失败!")
		}
		entity := res.AccessToken{User: user, Token: token, ExpiresAt: claims.StandardClaims.ExpiresAt * 1000}
		return &entity, nil
	}
}
