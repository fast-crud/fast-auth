package auth

import "github.com/golang-jwt/jwt"

type Claims struct {
	Id         uint
	Username   string
	RoleIds    []uint
	BufferTime int64
	Type       string //认证类型：user,app
	jwt.StandardClaims
}
