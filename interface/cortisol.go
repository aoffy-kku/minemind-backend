package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type CortisolServiceInterface interface {
	CreateCortisol(req model.CreateCortisolRequestJSON) (*model.CortisolJSON, error)
	CreateMultipleCortisol(req model.CreateMultipleCortisol) ([]*model.CortisolJSON, error)
	ToJSON(cortisol *model.Cortisol) *model.CortisolJSON
}
