package auth

import (
	"gin-sample-framework/config"
	"gin-sample-framework/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generate Token String
func GenerateTokenString(user_id int, user_role string) (string, error) {
	tokenModel := jwt.New(jwt.SigningMethodHS256)
	claims := tokenModel.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["user_role"] = user_role
	exp := time.Duration(config.Configuration.Jwt.Expire)
	claims["exp"] = utils.GetUTCTime().Add(time.Hour * exp).Unix()
	token, err := tokenModel.SignedString([]byte(config.Configuration.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
