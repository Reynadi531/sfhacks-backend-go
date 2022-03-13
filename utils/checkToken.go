package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func CheckIfTokenValid(at string) (bool, error) {
	var err error
	token, err := jwt.Parse(at, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_SECRET), nil
	})

	if token != nil && err == nil && token.Valid {
		return true, nil
	}

	return false, err
}

func ExtractClaims(at string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(at, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
