package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cortisol struct {
	Id primitive.ObjectID `bson:"_id"`
	UserId string         `bson:"user_id"`
	Cortisol float64      `bson:"cortisol"`
	Timestamp int64       `bson:"timestamp"`
	CreatedAt time.Time   `bson:"created_at"`
	CreatedBy string      `bson:"created_by"`
}

type CortisolJSON struct {
	Id primitive.ObjectID `json:"id"`
	UserId string         `json:"userId"`
	Cortisol float64      `json:"cortisol"`
	Timestamp int64       `json:"timestamp"`
	CreatedAt time.Time   `json:"createdAt"`
	CreatedBy string      `json:"-"`
}

type CreateCortisolRequestJSON struct {
	Cortisol float64   `json:"cortisol" validate:"required"`
	Timestamp int64 `json:"timestamp" validate:"required"`
	UserId string `json:"-"`
}

type CreateMultipleCortisol struct {
	Data []CreateCortisolRequestJSON `json:"data" validate:"gt=0,required"`
	UserId string `json:"-"`
}