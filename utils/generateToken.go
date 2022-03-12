package utils

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	JWT_SECRET = os.Getenv("JWT_SECRET")
	JWT_EXPIRE = time.Now().Add(time.Hour * 36).Unix()
)

func GenerateToken(id string) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["userid"] = id
	atClaims["exp"] = JWT_EXPIRE
	atClaims["authorized"] = true

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return token, nil
}
