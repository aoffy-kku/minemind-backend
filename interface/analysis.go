package _interface

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"time"
)

type AnalysisServiceInterface interface {
	CreateAnalysis(req model.CreateAnalysisRequestJSON) error
	GetAnalysisByDate(id string, date time.Time) ([]*model.AnalysisJSON, error)
	GetAnalysisByMode(id string, mode ...int64) ([]*model.AnalysisJSON, error)
	ToJSON(aggregate *model.Analysis) *model.AnalysisJSON
}
