package response

import (
	"gin-sample-framework/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`    // 业务码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据
}

type ResponseHandler struct{}

var Handler = &ResponseHandler{}

func (r *ResponseHandler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func (r *ResponseHandler) Error(c *gin.Context, err error) {
	var code int
	var message string

	if e, ok := err.(*errors.Error); ok {
		code = e.Code()
		message = e.Error()
	} else if e, ok := err.(errors.ErrorCode); ok {
		code = e.Code
		message = e.Message
	} else {
		code = errors.InternalServerError.Code
		message = err.Error()
	}

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func (r *ResponseHandler) BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    errors.BadRequest.Code,
		Message: err.Error(),
		Data:    nil,
	})
}

func (r *ResponseHandler) Custom(c *gin.Context, httpStatus int, code int, message string, data interface{}) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
