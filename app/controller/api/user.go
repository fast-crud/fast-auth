package api

import (
	"context"
	"github.com/fast-crud/fast-auth/app/controller/util"
	model "github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/frame/g"
)

type UserController struct{}

// MeReq -----------------------------------------------------
type MeReq struct {
	g.Meta `path:"/me" method:"post" per:"false"`
}
type MeRes struct {
	model.User
}

func (UserController) Me(ctx context.Context, req *MeReq) (res *MeRes, err error) {
	var request = g.RequestFromCtx(ctx)
	var id = util.Auth.GetUserInfo(request).Id
	data, err := system.UserService.GetById(id)
	if err != nil {
		return nil, err
	}
	return &MeRes{User: *data}, nil
}

// UpdateReq -----------------------------------------------------
type UpdateReq struct {
	g.Meta   `path:"/update" method:"post"`
	Avatar   string `json:"avatar" example:"用户头像"`
	NickName string `json:"nickName" example:"用户昵称"`
}
type UpdateRes struct {
	model.User
}

func (UserController) Update(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error) {
	var request = g.RequestFromCtx(ctx)
	var id = util.Auth.GetUserInfo(request).Id
	var info = system.UserUpdateParams{
		Id:       id,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}

	data, err := system.UserService.Update(&info)
	if err != nil {
		return nil, err
	}
	return &UpdateRes{*data}, nil
}

//
// ChangePasswordReq
// @Description:
//
type ChangePasswordReq struct {
	g.Meta      `path:"/changePassword" method:"post"`
	Password    string `json:"password" example:"密码"`
	NewPassword string `json:"newPassword" example:"新密码"`
}
type ChangePasswordRes struct {
}

//
// ChangePassword
// @Description:
// @receiver UserController
// @param ctx
// @param req
// @return *response.Response
//
func (UserController) ChangePassword(ctx context.Context, req *ChangePasswordReq) (res *ChangePasswordRes, err error) {

	var request = g.RequestFromCtx(ctx)
	var user = util.Auth.GetUserInfo(request)
	var id = user.Id
	var params = system.UserChangePasswordParams{
		Id:          id,
		Password:    req.Password,
		NewPassword: req.NewPassword,
	}
	if err := system.UserService.ChangePassword(&params); err != nil {
		return nil, err
	}
	return &ChangePasswordRes{}, nil
}
