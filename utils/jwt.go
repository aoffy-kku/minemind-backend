package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type (
	JWTConfig struct {
		Skipper    middleware.Skipper
		SigningKey interface{}
	}
	Skipper      func(ctx echo.Context) bool
	JWTExtractor func(ctx echo.Context) (string, error)
)

func GenerateAccessToken(id string, exp int64) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = exp

	t, _ := token.SignedString([]byte("Minemind2019"))
	return t
}

func GetClaims(tk string, config JWTConfig) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.SigningKey, nil
	})
	if err != nil {
		return nil, EchoHttpResponse(nil, http.StatusUnauthorized, HttpResponse{})
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, EchoHttpResponse(nil, http.StatusForbidden, HttpResponse{})
}

