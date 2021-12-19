package auth

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	Id         uint
	Username   string
	RoleIds    []uint
	BufferTime int64
	jwt.StandardClaims
}
