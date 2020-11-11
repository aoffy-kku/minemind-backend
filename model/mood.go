package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Mood struct {
	Id primitive.ObjectID
	Name string
	CreatedAt time.Time
	CreatedBy string
}

type CreateMoodRequestJSON struct {
	Name string
}