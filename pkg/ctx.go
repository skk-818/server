package pkg

import (
	"github.com/gin-gonic/gin"
	"server/pkg/jwt"
)

func GetClaims(c *gin.Context) *jwt.CustomClaims {
	val, exists := c.Get("claims")
	if !exists {
		return nil
	}
	claims, ok := val.(*jwt.CustomClaims)
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
