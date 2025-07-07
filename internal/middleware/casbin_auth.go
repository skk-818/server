package middleware

import (
	"server/pkg/response"
	"server/pkg/xerror"

	"github.com/gin-gonic/gin"
)

type CabinEnforce interface {
	Enforce(sub, obj, act string) (ok bool, err error)
}

type CasbinMiddleware struct {
	CabinEnforce
}

func NewCasbinMiddleware(cabinEnforce CabinEnforce) *CasbinMiddleware {
	return &CasbinMiddleware{
		cabinEnforce,
	}
}

func (m *CasbinMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleKeysVal, ok := c.Get("userRoles")
		if !ok {
			response.Fail(c, xerror.ErrUnauthorized) // 没有角色信息
			c.Abort()
			return
		}

		roleKeys, ok := roleKeysVal.([]string)
		if !ok || len(roleKeys) == 0 {
			response.Fail(c, xerror.ErrUnauthorized)
			c.Abort()
			return
		}

		obj := c.Request.URL.Path
		act := c.Request.Method

		// 遍历多个角色，任意一个角色拥有权限即放行
		for _, role := range roleKeys {
			ok, err := m.Enforce(role, obj, act)
			if err != nil {
				response.Fail(c, xerror.ErrInternal)
				c.Abort()
				return
			}
			if ok {
				c.Next()
				return
			}
		}

		// 所有角色都没有权限
		response.Fail(c, xerror.ErrPermissionDenied)
		c.Abort()
	}
}
