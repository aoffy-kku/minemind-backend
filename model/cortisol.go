package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cortisol struct {
	Id primitive.ObjectID
	UserId string
	Value float64
	Timestamp int64
	CreatedAt time.Time
	CreatedBy string
}

type CreateCortisolRequestJSON struct {
	Value float64
	Timestamp int64
}
type CreateMultipleCortisol struct {
	Data []CreateCortisolRequestJSON `json:"data" validate:"gt=0,required"`
}