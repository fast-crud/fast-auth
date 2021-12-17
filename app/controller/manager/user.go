package manager

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/app/model/basic/res"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/auth"
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

func (UserController) Add(ctx context.Context, req *RegisterReq) (res.CommonRes, error) {
	var info = request.UserRegister{
		Avatar:   req.Avatar,
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
	}
	data, err := system.UserService.Register(&info)
	if err != nil {
		return res.Error("注册失败！", err), err
	}
	return res.Success(g.Map{"user": data}), nil
}

type GetReq struct {
	g.Meta `path:"/get" method:"post"`
	Id     uint `json:"id" example:"id"`
}

func (UserController) Get(ctx context.Context, req *GetReq) (res.CommonRes, error) {
	data, err := system.UserService.GetById(req.Id)
	if err != nil {
		return res.Error("获取用户信息失败！", err), err
	}
	return res.Success(g.Map{"userInfo": data}), nil
}

type UpdateReq struct {
	g.Meta   `path:"/update" method:"post"`
	Avatar   string `json:"avatar" example:"用户头像"`
	NickName string `json:"nickName" example:"用户昵称"`
}

func (UserController) Update(ctx context.Context, req *UpdateReq) (res.CommonRes, error) {
	var request = g.RequestFromCtx(ctx)
	var id = auth.Claims.GetUserInfo(request).Id
	var info = system.UserUpdateParams{
		Id:       id,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}

	data, err := system.UserService.Update(&info)
	if err != nil {
		return res.Error("更新失败！", err), err
	}
	return res.Success(data), nil
}

type ResetPasswordReq struct {
	g.Meta      `path:"/resetPassword" method:"post"`
	Id          uint   `json:"id" example:"id"`
	NewPassword string `json:"newPassword" example:"新密码"`
}

func (UserController) ResetPassword(ctx context.Context, req *ResetPasswordReq) (res.CommonRes, error) {
	if err := system.UserService.ResetPassword(req.Id, req.NewPassword); err != nil {
		return res.Error("修改失败！", err), err
	}
	return res.Success(nil), nil
}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"post"`
	Id     uint `json:"id" example:"id"`
}

func (UserController) Delete(ctx context.Context, req *DeleteReq) (res.CommonRes, error) {
	if req.Id == 0 {
		return res.Error("id不能为空", nil), errors.New("id不能为空")
	}
	if err := system.UserService.Delete(req.Id); err != nil {
		return res.Error("删除失败", err), err
	}
	return res.Success(nil), nil
}
