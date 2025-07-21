package jwtx

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID   uint
	Username string
	Roles    []string
	jwt.RegisteredClaims
}

func (cl *CustomClaims) GetUserID() uint {
	return cl.UserID
}
