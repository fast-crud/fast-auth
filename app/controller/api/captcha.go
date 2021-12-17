package api

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/app/model/basic/res"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaController struct{}

type GenerateReq struct {
	g.Meta `path:"/" method:"post"`
}

func (UserController) Generate(ctx context.Context, req *GenerateReq) (res.CommonRes, error) {
	data, err := system.Captcha.Captcha()
	if err != nil {
		return res.Error("验证码获取失败!", err), err
	}
	return res.Success(data), nil
}
