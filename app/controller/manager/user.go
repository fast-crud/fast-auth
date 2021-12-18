package manager

import (
	"context"
	"github.com/fast-crud/fast-auth/app/controller/util"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	model "github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

type UserController struct{}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}
type RegisterRes struct {
	Id uint
}

func (UserController) Add(ctx context.Context, req *RegisterReq) (*RegisterRes, error) {
	var info = system.UserRegisterParams{
		Avatar:   req.Avatar,
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
	}
	data, err := system.UserService.Register(&info)
	if err != nil {
		return nil, err
	}
	return &RegisterRes{Id: data.Id}, nil
}

type GetReq struct {
	g.Meta `path:"/get" method:"post"`
	Id     uint `json:"id" example:"id"`
}
type GetRes struct {
	model.User
}

func (UserController) Get(ctx context.Context, req *GetReq) (res *GetRes, err error) {
	data, err := system.UserService.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return &GetRes{*data}, nil
}

type UpdateReq struct {
	g.Meta   `path:"/update" method:"post"`
	Avatar   string `json:"avatar" example:"用户头像"`
	NickName string `json:"nickName" example:"用户昵称"`
}

func (UserController) Update(ctx context.Context, req *UpdateReq) (*res.BlankRes, error) {
	var request = g.RequestFromCtx(ctx)
	var id = util.Claims.GetUserInfo(request).Id
	var info = system.UserUpdateParams{
		Id:       id,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}

	_, err := system.UserService.Update(&info)
	if err != nil {
		return nil, err
	}
	return &res.BlankRes{}, nil
}

type ResetPasswordReq struct {
	g.Meta      `path:"/resetPassword" method:"post"`
	Id          uint   `json:"id" example:"id"`
	NewPassword string `json:"newPassword" example:"新密码"`
}

func (UserController) ResetPassword(ctx context.Context, req *ResetPasswordReq) (*res.BlankRes, error) {
	if err := system.UserService.ResetPassword(req.Id, req.NewPassword); err != nil {
		return nil, err
	}
	return res.NewBlankRes(), nil
}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"post"`
	Id     uint `json:"id" example:"id"`
}

func (UserController) Delete(ctx context.Context, req *DeleteReq) (*res.BlankRes, error) {
	if req.Id == 0 {
		return nil, errors.New("id不能为空")
	}
	if err := system.UserService.Delete(req.Id); err != nil {
		return nil, err
	}
	return res.NewBlankRes(), nil
}
