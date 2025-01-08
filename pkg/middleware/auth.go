package middleware

import (
	"fmt"
	"gin-sample-framework/config"
	"gin-sample-framework/errors"
	"gin-sample-framework/pkg/auth"
	"gin-sample-framework/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, H(errors.InvalidToken, nil, errors.InvalidToken.String()))
			return
		}

		tokenID := extractSessionID(authHeader)
		if tokenID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, H(errors.InvalidToken, nil, errors.InvalidToken.String()))
			return
		}

		token, err := jwt.Parse(tokenID, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Configuration.Jwt.Secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, H(errors.InvalidToken, nil, errors.InvalidToken.String()))
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		user_id := claims["user_id"]
		user_role := claims["user_role"]
		expInterface := claims["exp"]
		if exp := time.Unix(int64(expInterface.(float64)), 0); exp.Before(utils.GetUTCTime()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, H(errors.InvalidToken, nil, errors.InvalidToken.String()))
			return
		}
		if user_id == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, H(errors.InvalidToken, nil, errors.InvalidToken.String()))
			return
		}
		var (
			out auth.Token
		)
		if user_role != nil {
			out.UserRole = user_role.(string)
		}
		out.UserId = int(user_id.(float64))
		c.Request = c.Request.WithContext(auth.SetTokenData(c.Request.Context(), out))
		c.Next()
	}
}

func extractSessionID(authHeader string) string {
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return strings.TrimSpace(splitToken[1])
}
