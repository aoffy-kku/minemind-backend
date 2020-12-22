package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type UserEvaluationServiceInterface interface {
	CreateUserEvaluation(req model.CreateUserEvaluationRequestJSON) (*model.UserEvaluationJSON, error)
	GetLatestUserEvaluation(id string) ([]*model.UserEvaluationJSON, error)
	ToJSON(evaluation *model.UserEvaluation) *model.UserEvaluationJSON
}
