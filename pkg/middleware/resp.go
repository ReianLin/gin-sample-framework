package middleware

import (
	"bytes"
	"gin-sample-framework/errors"
	"gin-sample-framework/pkg/logger"

	"net/http"
	"runtime/debug"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func HandleResponse(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			err error
		)

		defer func() {
			if r := recover(); r != nil {
				// 在这里处理系统崩溃信息，可以记录日志、发送警报等
				logger.Errorf("系统错误%s,%s", r, string(debug.Stack()))
				// 返回 500 Internal Server Error
				c.JSON(http.StatusInternalServerError, H(errors.InternalServerError, nil, "服务器内部错误"))
			}
		}()

		c.Request = c.Request.WithContext(c.Request.Context())
		newWriter := customWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = newWriter

		c.Next()

		//错误处理
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				err = e.Err
				switch errType := err.(type) {
				case errors.CustomError: //判断是否业务错误
					getType := errors.GetType(errType)
					switch getType {
					case errors.InvalidToken:
						c.JSON(http.StatusUnauthorized, nil)
					default:
						c.JSON(http.StatusOK, H(getType, nil, getType.String()))
					}
				case validator.ValidationErrors: //判断是否POST入参校验错误
					c.JSON(http.StatusBadRequest, H(errors.EmptyParameter, nil, errors.EmptyParameter.String()))
				default:
					c.JSON(http.StatusOK, H(errors.InternalServerError, nil, errors.InternalServerError.String()))
				}
			}
		}
	}
}

type customWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (c customWriter) Write(p []byte) (int, error) {
	if _, err := c.ResponseWriter.Write(p); err != nil {
		return 0, err
	}
	return c.body.Write(p)
}

func H(status errors.Type, data any, msg string) gin.H {
	return gin.H{
		"code": status,
		"data": data,
		"msg":  msg,
	}
}
