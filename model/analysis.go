package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Analysis struct {
	Id primitive.ObjectID                  `bson:"_id"`
	Mode int64                          `bson:"mode"`
	UserId string                          `bson:"user_id"`
	UserEvaluationId primitive.ObjectID `bson:"user_evaluation_id"`
	UserEvaluation UserEvaluation       `bson:"user_evaluation"`
	CortisolId primitive.ObjectID          `bson:"cortisol_id"`
	Cortisol Cortisol                      `bson:"cortisol"`
	Status int64                           `bson:"status"`
	Class int64                   `bson:"class"`
	Score float64                 `bson:"score"`
	CreatedAt time.Time                    `bson:"created_at"`
	CreatedBy string                       `bson:"created_by"`
	UpdatedAt time.Time                    `bson:"updated_at"`
	UpdatedBy string                       `bson:"updated_by"`
}

type AnalysisJSON struct {
	Id primitive.ObjectID                  `json:"id"`
	Mode int64                          `json:"mode"`
	UserId string                          `json:"userId"`
	UserEvaluationId primitive.ObjectID `json:"userEvaluationId"`
	UserEvaluation UserEvaluationJSON   `json:"userEvaluation"`
	CortisolId primitive.ObjectID          `json:"cortisolId"`
	Cortisol CortisolJSON                  `json:"cortisol"`
	Status int64                           `json:"status"`
	Class int64                   `json:"class"`
	Score float64                 `json:"score"`
	CreatedAt time.Time                    `json:"createdAt"`
	CreatedBy string                       `json:"createdBy"`
	UpdatedAt time.Time                    `json:"updatedAt"`
	UpdatedBy string                       `json:"updatedBy"`
}

type CreateAnalysisRequestJSON struct {
	Mode int64 `json:"mode" validate:"required"`
	UserId string `json:"-"`
}

type UpdateAnalysisRequestJSON struct {
	Id primitive.ObjectID `json:"-"`
	Class int64                   `json:"class" validate:"required"`
	Score float64                 `json:"score" validate:"required"`
}