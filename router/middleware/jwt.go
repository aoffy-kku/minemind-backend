package middleware

import (
	"context"
	"fmt"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func JWT(key interface{}, db *mongo.Database) echo.MiddlewareFunc {
	c := utils.JWTConfig{}
	c.SigningKey = key
	return JWTWithConfig(c, db)
}

func JWTWithConfig(config utils.JWTConfig, db *mongo.Database) echo.MiddlewareFunc {
	ctx := context.Background()
	extractor := jwtFromHeader("Authorization", "Bearer")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractor(c)
			if err != nil {
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return utils.EchoHttpResponse(c, http.StatusUnauthorized, utils.HttpResponse{})
			}
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return utils.EchoHttpResponse(c, http.StatusForbidden, utils.HttpResponse{Message: err.Error()})
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				id := claims["id"].(string)
				var m interface{}
				if err := db.Collection("access_token").FindOne(ctx, bson.M{
					"user_id": bson.M{
						"$eq": id,
					},
					"pair": bson.M{
						"$eq": token.Raw,
					},
				}).Decode(&m); err != nil {
					return utils.EchoHttpResponse(c, http.StatusUnauthorized, utils.HttpResponse{})
				}
				c.Set("id", id)
				c.Set("access_token", token.Raw)
				return next(c)
			}
			return utils.EchoHttpResponse(c, http.StatusForbidden, utils.HttpResponse{})
		}
	}
}

func jwtFromHeader(header string, authScheme string) utils.JWTExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", fmt.Errorf(http.StatusText(401))
	}
}
