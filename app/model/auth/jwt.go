package auth

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	Id         uint
	NickName   string
	Username   string
	RoleIds    []uint
	BufferTime int64
	jwt.StandardClaims
}
