package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/xerror"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

const SuccessCode = 200
const SuccessMessage = "success"

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, Response[any]{
		Code:    SuccessCode,
		Message: SuccessMessage,
	})
}

func SuccessWithData[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Code:    SuccessCode,
		Message: SuccessMessage,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	if bizErr, ok := err.(*xerror.BizError); ok {
		c.JSON(http.StatusOK, Response[any]{
			Code:    bizErr.Code,
			Message: bizErr.Message,
		})
	} else {
		c.JSON(http.StatusOK, Response[any]{
			Code:    xerror.ErrInternalServer.Code,
			Message: "系统异常：" + err.Error(),
		})
	}
	c.Abort()
}
