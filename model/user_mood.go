package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserMood struct {
	Id primitive.ObjectID `bson:"_id"`
	Name string           `bson:"name"`
	UserId string         `bson:"user_id"`
	Active bool           `bson:"active"`
	CreatedAt time.Time   `bson:"created_at"`
	CreatedBy string      `bson:"created_by"`
}

type UserMoodJSON struct {
	Id primitive.ObjectID `json:"id"`
	Name string           `json:"name"`
	UserId string         `json:"userId"`
	Active bool           `json:"active"`
	CreatedAt time.Time   `json:"createdAt"`
	CreatedBy string      `json:"createdBy"`
}

type CreateUserMoodJSON struct {
	Moods []string           `json:"name" validate:"required"`
	UserId string         		`json:"-"`
}

type UpdateUserMoodJSON struct {
	Id primitive.ObjectID `json:"-"`
	Name string           `json:"name" validate:"required"`
	Active bool           `json:"active"`
}