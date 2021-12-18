package api

import (
	"context"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type AuthController struct{}

type CallbackReq struct {
	g.Meta `path:"/callback" method:"get" auth:"false"`
}
type CallbackRes struct {
	Id          string    `json:"id"`
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expires"`
}

func (AuthController *AuthController) callback(ctx context.Context, req *CallbackReq) (*CallbackRes, error) {
	var r = g.RequestFromCtx(ctx)
	var code = r.GetQuery("code").String()
	var state = r.GetQuery("state").String()
	token, err := auth.GetOAuthToken(code, state)
	if err != nil {
		panic(err)
	}

	claims, err := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		panic(err)
	}

	claims.AccessToken = token.AccessToken

	return &CallbackRes{claims.User.Id, claims.AccessToken, claims.RegisteredClaims.ExpiresAt.Time}, nil
}

type GetLoginUrlReq struct {
	g.Meta `path:"/getLoginUrl" method:"get" auth:"false"`
}
type GetLoginUrlRes struct {
	Url string
}

func (AuthController *AuthController) GetLoginUrl(ctx context.Context, req *GetLoginUrlReq) (*GetLoginUrlRes, error) {
	var url = auth.GetSigninUrl("http://localhost:8199/api/auth/callback")

	return &GetLoginUrlRes{url}, nil
}
