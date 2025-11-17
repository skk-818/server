package pkg

import (
	"server/pkg/jwtx"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) *jwtx.CustomClaims {
	val, exists := c.Get("claims")
	if !exists {
		return nil
	}
	claims, ok := val.(*jwtx.CustomClaims)
	if !ok {
		return nil
	}
	return claims
}

func GetUserID(c *gin.Context) uint {
	if claims := GetClaims(c); claims != nil {
		return claims.GetUserID()
	}
	return 0
}

func GetRoles(c *gin.Context) []string {
	if claims := GetClaims(c); claims != nil {
		return claims.Roles
	}
	return nil
}
