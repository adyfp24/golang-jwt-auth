package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string)(*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("jwtsecretkey"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid{
		return nil, errors.New("invalid token")
	}

	return token, nil
}