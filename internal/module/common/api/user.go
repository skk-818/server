package api

import (
	"github.com/gin-gonic/gin"
	"server/pkg/response"
)

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

func (ua *UserApi) InitUserRouter(router *gin.RouterGroup) {
	router.POST("info", ua.Info)
	router.POST("register", ua.Info)
}

func (ua *UserApi) Info(c *gin.Context) {
	response.Success(c)

}

func (ua *UserApi) Register(c *gin.Context) {
	response.Success(c)
}
