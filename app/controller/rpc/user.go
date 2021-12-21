package rpc

import (
	"context"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

type UserController struct {
	g.Meta `path:"/user"`
}

type ResetPasswordReq struct {
	g.Meta      `path:"/resetPassword" method:"post" auth:"false"`
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
	g.Meta `path:"/delete" method:"post" auth:"false"`
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
