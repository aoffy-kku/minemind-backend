package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type UserEvaluationResult struct {
	QuestionId primitive.ObjectID
	OptionId primitive.ObjectID
}

type UserEvaluation struct {
	Id primitive.ObjectID
	UserId string
	EvaluationId string
	Result []UserEvaluationResult
	CreatedAt time.Time
	CreatedBy string
}

type UserEvaluationJSON struct {
	Id primitive.ObjectID
	UserId string
	EvaluationId string
	Result []UserEvaluationResult
	CreatedAt time.Time
	CreatedBy string
}

type CreateUserEvaluationRequestJSON struct {
	EvaluationId string
	Result []UserEvaluationResult
}
