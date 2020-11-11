package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Option struct {
	Id primitive.ObjectID
	QuestionId primitive.ObjectID
	Title string
	Value int64
}
type Question struct {
	Id primitive.ObjectID
	EvaluationId string
	Title string
	Options []primitive.ObjectID
}
type Evaluation struct {
	Id string
	Description string
	Questions []primitive.ObjectID
	CreatedAt time.Time
	CreatedBy string
}
type CreateEvaluationRequestJSON struct {
	Name string
	Description string
	Questions []Question
}
