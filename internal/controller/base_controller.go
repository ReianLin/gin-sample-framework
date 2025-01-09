package controller

import (
	"gin-sample-framework/pkg/response"

	"github.com/gin-gonic/gin"
)

type baseController struct {
}

func (b *baseController) Success(c *gin.Context, data interface{}) {
	response.Handler.Success(c, data)
}

func (b *baseController) Error(c *gin.Context, err error) {
	response.Handler.Error(c, err)
}

func (b *baseController) BadRequest(c *gin.Context, err error) {
	response.Handler.BadRequest(c, err)
}

func (b *baseController) Custom(c *gin.Context, httpStatus int, code int, message string, data interface{}) {
	response.Handler.Custom(c, httpStatus, code, message, data)
}
