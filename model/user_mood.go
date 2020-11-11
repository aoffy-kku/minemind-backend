package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserMood struct {
	Id primitive.ObjectID
	Name string
	UserId string
	Active bool
	CreatedAt time.Time
	CreatedBy string
}

type CreateUserMoodRequest struct {
	Name string
}

type UserMoodResponse struct {
	Id primitive.ObjectID
	Name string
}