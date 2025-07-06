package jwt

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID   uint
	Username string
	Role     string
	jwt.RegisteredClaims
}

func (cl *CustomClaims) GetUserID() uint {
	return cl.UserID
}
