package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mfarooqzahid/go-social/internal/config"
	"github.com/mfarooqzahid/go-social/internal/models"
)

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(config.Envs.JWTSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
