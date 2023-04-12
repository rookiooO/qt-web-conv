package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	Unique int64
	jwt.RegisteredClaims
}

var secret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJzdWIiOiIwOTg3NjU0MzIxIiwibmFtZSI6InFpdXJlbnpoaSIsImlhdCI6MTUxNjIzOTAyMn0." +
	"-CwIcfolmA1dUtyOCGFkOr_BdaY595bpD4GL2NFKbCc")

func GenToken(unique int64) (string, error) {
	var (
		customClaims *CustomClaims
		token        *jwt.Token
		tokenString  string
		err          error
	)
	customClaims = &CustomClaims{
		Unique: unique,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(100) * time.Second)),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err = token.SignedString(secret)

	return tokenString, err
}
