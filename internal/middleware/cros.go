package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/core/config"
	"strings"
)

type corsOptions struct {
	allowOrigins []string
}

type CorsMiddleware struct {
	options *corsOptions
}

func NewCorsMiddleware(cfg *config.Cors) *CorsMiddleware {
	return &CorsMiddleware{options: &corsOptions{allowOrigins: cfg.AllowOrigins}}
}

func (cm *CorsMiddleware) Handler() gin.HandlerFunc {

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" || !matchOrigin(origin, cm.options.allowOrigins) {
			c.Next()
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Origin,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// 简单的通配符匹配，支持 *.xxx.com 或 IP 段
func matchOrigin(origin string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.HasPrefix(pattern, "*.") {
			// 匹配 *.domain.com
			base := strings.TrimPrefix(pattern, "*.")
			if strings.HasSuffix(origin, base) {
				return true
			}
		} else if strings.Contains(pattern, "*") {
			// IP 通配匹配，例如 192.168.*.*
			regex := strings.ReplaceAll(pattern, ".", "\\.")
			regex = strings.ReplaceAll(regex, "*", ".*")
			matched := strings.HasPrefix(origin, "http://") || strings.HasPrefix(origin, "https://")
			if matched && matchWildcard(origin, regex) {
				return true
			}
		} else if origin == pattern {
			return true
		}
	}
	return false
}

// 粗略正则通配支持
func matchWildcard(origin, pattern string) bool {
	return strings.Contains(origin, pattern[:len(pattern)-2]) // 这里只是简单实现，可替换为正则
}
