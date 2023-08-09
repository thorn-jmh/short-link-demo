package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint
	jwt.MapClaims
}

func Release(claim *Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}

func Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
