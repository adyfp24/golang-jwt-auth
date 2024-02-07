package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(userID uint) (string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("jwtsecretkey"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}