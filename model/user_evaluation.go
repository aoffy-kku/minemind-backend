package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserEvaluation struct {
	Id primitive.ObjectID `bson:"_id"`
	UserId string         `bson:"user_id"`
	EvaluationId string   `bson:"evaluation_id"` 
	Name string           `bson:"name"`
	Description string    `bson:"description"`
	Questions []Question  `bson:"questions"`
	CreatedAt time.Time   `bson:"created_at"`
	CreatedBy string      `bson:"created_by"`
}

type UserEvaluationJSON struct {
	Id primitive.ObjectID            `json:"id"`
	EvaluationId string `json:"evaluationId"`
	UserId string `json:"userId"`
	Name string          `json:"name"`
	Description string   `json:"description"`
	Questions []QuestionJSON `json:"questions"`
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
}

type CreateUserEvaluationRequestJSON struct {
	UserId string `json:"-"`
	EvaluationId string `json:"evaluationId"`
	Name string          `json:"name"`
	Description string   `json:"description"`
	Questions []Question `json:"questions"`
}