package middleware

import (
	"context"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

const (
	Admin   = "admin"
	Officer = "officer"
	User    = "user"
)

func Roles(db *mongo.Database, roles ...string) echo.MiddlewareFunc {
	ctx := context.Background()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, ok := c.Get("id").(string)
			if !ok {
				return utils.EchoHttpResponse(c, http.StatusUnauthorized, utils.HttpResponse{})
			}
			user := db.Collection("user")
			var m interface{}
			if err := user.FindOne(ctx, bson.M{
				"_id": bson.M{
					"$eq": id,
				},
				"roles": bson.M{
					"$in": roles,
				},
			}).Decode(&m); err != nil {
				return utils.EchoHttpResponse(c, http.StatusMethodNotAllowed, utils.HttpResponse{})
			}
			return next(c)
		}
	}
}

