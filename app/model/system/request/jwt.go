package request

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Id         uint
	NickName   string
	Username   string
	RoleIds    []uint
	BufferTime int64
	jwt.StandardClaims
}
