package api

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

type CaptchaController struct{}

type GenerateReq struct {
	g.Meta `path:"/" method:"post"`
}
type GenerateRes struct {
}

func (UserController) Generate(ctx context.Context, req *GenerateReq) (*system.CaptchaRes, error) {
	data, err := system.Captcha.Captcha()
	if err != nil {
		return nil, errors.Wrap(err, "验证码生成失败")
	}
	return data, nil
}
