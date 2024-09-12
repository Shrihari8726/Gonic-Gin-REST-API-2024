// utils/jwt.go
package utils

import (
	"time"

	"example.com/gonicginrestapi/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email string) (string, error) {
	secret := config.GetJWTSecret()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := config.GetJWTSecret()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return token, err
}
