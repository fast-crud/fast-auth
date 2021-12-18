package basic

import (
	"github.com/fast-crud/fast-auth/app/model/auth"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type JwtService struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JwtService {
	return &JwtService{[]byte(global.Config.Jwt.SigningKey)}
}

// CreateToken 创建一个token
func (jwtService *JwtService) CreateToken(claims *auth.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtService.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (jwtService *JwtService) CreateTokenByOldToken(oldToken string, claims *auth.JwtClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JwtService:"+oldToken, func() (interface{}, error) {
		return jwtService.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (jwtService *JwtService) ParseToken(tokenString string) (*auth.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.JwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtService.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*auth.JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
