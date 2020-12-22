package service

import (
	"context"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserEvaluationService struct {
	db  *mongo.Database
	col *mongo.Collection
}

func (u *UserEvaluationService) CreateUserEvaluation(req model.CreateUserEvaluationRequestJSON) (*model.UserEvaluationJSON, error) {
	ctx := context.Background()
	now := time.Now()
	result, err := u.col.InsertOne(ctx, model.UserEvaluation{
		Id:           primitive.NewObjectIDFromTimestamp(now),
		UserId:       req.UserId,
		EvaluationId: req.EvaluationId,
		Name:         req.Name,
		Description:  req.Description,
		Questions:    req.Questions,
		CreatedAt:    now,
		CreatedBy:    req.UserId,
	})
	if err != nil {
		return nil, err
	}
	var m *model.UserEvaluation
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": result.InsertedID,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToJSON(m), nil
}

func (u *UserEvaluationService) GetLatestUserEvaluation(id string) ([]*model.UserEvaluationJSON, error) {
	ctx := context.Background()
	opts := &options.FindOptions{}
	opts.SetLimit(1)
	opts.SetSort(bson.M{
		"_id": -1,
	})
	var docs []*model.UserEvaluationJSON
	cur, err := u.col.Find(ctx, bson.M{
		"user_id": bson.M{
			"$eq": id,
		},
		"evaluation_id": bson.M{
			"$eq": "st5",
		},
	}, opts)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.UserEvaluation
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		docs = append(docs, u.ToJSON(&m))
	}

	cur, err = u.col.Find(ctx, bson.M{
		"user_id": bson.M{
			"$eq": id,
		},
		"evaluation_id": bson.M{
			"$eq": "phq9",
		},
	}, opts)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.UserEvaluation
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		docs = append(docs, u.ToJSON(&m))
	}
	return docs, nil
}

func (u *UserEvaluationService) ToJSON(evaluation *model.UserEvaluation) *model.UserEvaluationJSON {
	var questions []model.QuestionJSON
	for _, q := range evaluation.Questions {
		var opts []model.OptionJSON
		for _, o := range q.Options {
			opts = append(opts, model.OptionJSON{
				Id:    o.Id,
				Title: o.Title,
				Value: o.Value,
			})
		}
		questions = append(questions, model.QuestionJSON{
			Id:      q.Id,
			Title:   q.Title,
			Options: opts,
		})
	}
	return &model.UserEvaluationJSON{
		Id:          evaluation.Id,
		Name:        evaluation.Name,
		UserId:      evaluation.UserId,
		Description: evaluation.Description,
		EvaluationId: evaluation.EvaluationId,
		Questions:   questions,
		CreatedAt:   evaluation.CreatedAt,
		CreatedBy:   evaluation.CreatedBy,
	}
}

func NewUserEvaluationService(db *mongo.Database) *UserEvaluationService {
	return &UserEvaluationService{
		db:  db,
		col: db.Collection("user_evaluation"),
	}
}
