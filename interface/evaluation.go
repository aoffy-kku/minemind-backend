package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type EvaluationServiceInterface interface {
	GetEvaluations() ([]*model.EvaluationJSON, error)
	ToJSON(evaluation *model.Evaluation) *model.EvaluationJSON
}
