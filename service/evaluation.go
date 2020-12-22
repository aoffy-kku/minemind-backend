package service

import (
	"context"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EvaluationService struct {
	db *mongo.Database
	col *mongo.Collection
}

func (e *EvaluationService) GetEvaluations() ([]*model.EvaluationJSON, error) {
	ctx := context.Background()
	var docs []*model.EvaluationJSON
	cur, err := e.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.Evaluation
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		docs = append(docs, e.ToJSON(&m))
	}
	return docs, nil
}

func (e *EvaluationService) ToJSON(evaluation *model.Evaluation) *model.EvaluationJSON {
	var questions []model.QuestionJSON
	for _, q := range evaluation.Questions {
		var options []model.OptionJSON
		for _, o := range q.Options {
			options = append(options, model.OptionJSON{
				Id:    o.Id,
				Title: o.Title,
				Value: o.Value,
			})
		}
		questions = append(questions, model.QuestionJSON{
			Id:      q.Id,
			Title:   q.Title,
			Options: options,
		})
	}
	return &model.EvaluationJSON{
		Id:          evaluation.Id,
		Name:        evaluation.Name,
		Description: evaluation.Description,
		Questions:   questions,
		CreatedAt:   evaluation.CreatedAt,
		CreatedBy:   evaluation.CreatedBy,
	}
}

func NewEvaluationService(db * mongo.Database) *EvaluationService {
	return &EvaluationService{
		db:  db,
		col: db.Collection("evaluation"),
	}
}
