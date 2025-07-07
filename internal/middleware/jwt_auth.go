package middleware

import (
	"server/pkg/jwt"
	"server/pkg/response"
	"server/pkg/xerror"
	"strings"

	"github.com/gin-gonic/gin"
)

type JwtParse interface {
	Parse(string string) (*jwt.CustomClaims, error)
}

type JwtMiddleware struct {
	JwtParse
}

func NewJwtMiddleware(parse JwtParse) *JwtMiddleware {
	return &JwtMiddleware{
		parse,
	}
}

func (jm *JwtMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取 Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, xerror.ErrAuthHeaderMissing)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && strings.EqualFold(parts[0], "Bearer")) {
			response.Fail(c, xerror.ErrAuthHeaderFormat)
			c.Abort()
			return
		}

		// 解析 JWT
		claims, err := jm.Parse(parts[1])
		if err != nil {
			response.Fail(c, xerror.ErrInvalidToken)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("userRoles", claims.Roles)

		c.Next()
	}
}
