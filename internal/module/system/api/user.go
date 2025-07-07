package api

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/config"
	"server/internal/module/system/service"
	"server/pkg/response"
)

type UserApi struct {
	config      *config.Config
	userService *service.UserService
}

func NewUserApi(config *config.Config, userService *service.UserService) *UserApi {
	return &UserApi{
		config: config,
	}
}

func (ua *UserApi) InitUserApi(router *gin.RouterGroup) {
	router.POST("info", ua.Info)
}

func (ua *UserApi) Info(c *gin.Context) {
	response.SuccessWithData(c, gin.H{"name": "", "roles": []string{"R_ADMIN"}, "introduction": "", "avatar": "", "email": ""})
}
