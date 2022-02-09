package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var sercretKey = "love"

func GenerateJwt(Issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    Issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte(sercretKey))
}

func ParseJwt(cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sercretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}
	return token.Claims.(*jwt.StandardClaims).Issuer, nil

}
