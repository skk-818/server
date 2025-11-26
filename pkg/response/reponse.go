package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/errorx"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

const SuccessCode = 200
const SuccessMessage = "success"

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, Response[any]{
		Code: SuccessCode,
		Msg:  SuccessMessage,
	})
}

func SuccessWithData[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Code: SuccessCode,
		Msg:  SuccessMessage,
		Data: data,
	})
}

func Fail(c *gin.Context, err error) {
	if bizErr, ok := err.(*errorx.BizError); ok {
		c.JSON(http.StatusOK, Response[any]{
			Code: bizErr.Code,
			Msg:  bizErr.Message,
		})
	} else {
		c.JSON(http.StatusOK, Response[any]{
			Code: errorx.ErrInternalServer.Code,
			Msg:  "系统异常：" + err.Error(),
		})
	}
	c.Abort()
}
