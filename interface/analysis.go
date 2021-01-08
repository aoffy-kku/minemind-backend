package _interface

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AnalysisServiceInterface interface {
	CreateAnalysis(req model.CreateAnalysisRequestJSON) error
	GetAnalysisByDate(id string, date time.Time) ([]*model.AnalysisJSON, error)
	GetAnalysisByMode(id string, mode ...int64) ([]*model.AnalysisJSON, error)
	DeleteAnalysis(id primitive.ObjectID, uid string) error
	ToJSON(aggregate *model.Analysis) *model.AnalysisJSON
}
