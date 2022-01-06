package system

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/common"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var UserService = new(userService)

type userService struct {
	g.Meta `path:"/user"`
}

type UserRegisterReq struct {
	g.Meta   `path:"/register" method:"post" auth:"false"`
	Avatar   string
	Username string
	Password string
	NickName string
}

type UserInfoRes struct {
	system.User
}

// Register
// @Description: 用户注册
// @receiver userService
// @param info
// @return data
// @return err
func (userService *userService) Register(req *UserRegisterReq) (res *UserInfoRes, err error) {
	var entity system.User
	if !errors.Is(global.Db.Where("username = ?", req.Username).First(&entity).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return nil, errors.Wrap(err, "用户名已注册")
	}
	entity = system.User{
		Avatar:   req.Avatar,
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
	}
	if err = entity.EncryptedPassword(); err != nil {
		return nil, errors.Wrap(err, "密码加密失败!")
	}
	if err = global.Db.Create(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "用户注册失败!")
	}
	return &UserInfoRes{User: entity}, nil
}

type UserGetByIdReq struct {
	g.Meta `path:"/getById" method:"post" auth:"false"`
	Id     uint
}

//
// GetById
// @Description: 根据id获取用户信息
// @receiver userService
// @param Id
// @return data
// @return err
//
func (userService *userService) GetById(req UserGetByIdReq) (res *UserInfoRes, err error) {
	if req.Id == 0 {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	var entity system.User
	var search = func(db *gorm.DB) *gorm.DB {
		db = db.Where("id = ?", req.Id)
		return db
	}
	if err = global.Db.Scopes(search).Preload("Roles").First(&entity).Error; err != nil {
		return nil, gerror.NewCode(constants.CodeUserFindError)
	}
	return &UserInfoRes{User: entity}, nil
}

type UserFindReq struct {
	g.Meta   `path:"/find" method:"post" auth:"false"`
	Id       uint   `json:"id" example:"7"`
	Username string `json:"username" example:"7"`
}
type UserListRes struct {
	List []system.User
}

//
//  Find
//  @Description: 查询user
//  @receiver userService
//  @param userFindParams
//  @return data
//  @return err
//
func (userService *userService) Find(req *UserFindReq) (res *UserListRes, err error) {
	var list []system.User
	if req.Username == "" && req.Id == 0 {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	if err = global.Db.Where(&req).Preload("Roles").Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "用户查询失败!")
	}
	return &UserListRes{List: list}, nil
}

type UserUpdateReq struct {
	g.Meta   `path:"/update" method:"post" auth:"false"`
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
func (userService *userService) Update(req *UserUpdateReq) (*res.BlankRes, error) {
	update := system.User{
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}
	var err = global.Db.Where("id = ?", req.Id).Updates(&update).Error
	return res.NewBlankRes(), err
}

type UserChangePasswordReq struct {
	g.Meta      `path:"/changePassword" method:"post" auth:"false"`
	Id          uint   `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// ChangePassword
// @Description: 修改密码
// @receiver userService
// @param params
// @return error
func (userService *userService) ChangePassword(req *UserChangePasswordReq) (*res.BlankRes, error) {
	var entity system.User
	err := global.Db.Where("id = ?", req.Id).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "用户不存在! ")
	}
	if !entity.CompareHashAndPassword(req.Password) {
		return nil, errors.Wrap(err, "密码错误!")
	}
	entity.Password = req.NewPassword
	if err = entity.EncryptedPassword(); err != nil {
		return nil, errors.Wrap(err, "密码加密失败!")
	}
	err = global.Db.Where("id = ?", req.Id).Update("password", entity.Password).Error
	return res.NewBlankRes(), err
}

type UserResetPasswordReq struct {
	g.Meta      `path:"/resetPassword" method:"post" auth:"false"`
	Id          uint   `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// RestPassword
// @Description: 重置密码
// @receiver userService
// @param params
// @return error
func (userService *userService) ResetPassword(req UserResetPasswordReq) (*res.BlankRes, error) {
	var entity system.User
	err := global.Db.Where("id = ?", req.Id).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "用户不存在! ")
	}
	entity.Password = req.NewPassword
	if err = entity.EncryptedPassword(); err != nil {
		return nil, errors.Wrap(err, "密码加密失败!")
	}
	err = global.Db.Model(&entity).Where("id = ?", req.Id).Update("password", entity.Password).Error
	return nil, err
}

type UserDeleteReq struct {
	g.Meta `path:"/resetPassword" method:"post" auth:"false"`
	Id     uint `json:"id"`
}

func (userService *userService) Delete(req UserDeleteReq) (*res.BlankRes, error) {
	var err = global.Db.Delete(&system.User{}, req.Id).Error
	return nil, err
}

type UserPageReq struct {
	g.Meta `path:"/page" method:"post" auth:"false"`
	common.PageInfo
	AppId          uint `json:"appId"`
	OrganizationId uint `json:"organizationId"`
}

type UserPageRes struct {
	common.PageInfo
	List []system.User
}

// GetList 获取用户列表
func (userService *userService) Page(req *UserPageReq) (*UserPageRes, error) {
	var entities []system.User
	var total int64
	db := global.Db.Model(&system.User{})
	var err = db.Count(&total).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询失败")
	}
	err = db.Scopes(common.Paginate(&req.PageInfo)).Preload("roleIds").Find(&entities).Error
	return &UserPageRes{List: entities, PageInfo: common.PageInfo{Total: total}}, err
}
