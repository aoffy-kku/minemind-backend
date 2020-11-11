package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Analysis struct {
	Id primitive.ObjectID
	Type string
	UserId string
	CortisolId primitive.ObjectID
	Status int64
	Result int64
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type CreateAnalysisRequestJSON struct {
	Type string
	CortisolId primitive.ObjectID
}

type UpdateAnalysisRequestJSON struct {
	Result int64
}