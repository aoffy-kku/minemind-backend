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

type UserDiaryMoodService struct {
	db *mongo.Database
	col *mongo.Collection
}

func (u *UserDiaryMoodService) GetUserDiaryMoodByDate(id string, date time.Time) ([]*model.UserDiaryMoodJSON, error) {
	ctx := context.Background()
	fromDate := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	toDate := fromDate.AddDate(0, 1, 0)
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{
		"_id": -1,
	})
	cur, err := u.col.Find(ctx, bson.M{
		"user_id": bson.M{
			"$eq": id,
		},
		"created_at": bson.M{
			"$gte": fromDate,
			"$lt": toDate,
		},
	}, opts)
	if err != nil {
		return nil, err
	}
	var docs []*model.UserDiaryMoodJSON
	for cur.Next(ctx) {
		var m model.UserDiaryMood
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		docs = append(docs, u.ToJSON(m))
	}
	return docs, nil
}

func (u *UserDiaryMoodService) CreateUserDiaryMood(request model.CreateUserDiaryMoodRequestJSON) (*model.UserDiaryMoodJSON, error) {
	ctx := context.Background()
	now := time.Now()
	result, err := u.col.InsertOne(ctx, &model.UserDiaryMood{
		Id:        primitive.NewObjectIDFromTimestamp(now),
		UserId:    request.UserId,
		Moods:     request.Moods,
		CreatedAt: now,
		CreatedBy: request.UserId,
	})
	if err != nil {
		return nil, err
	}
	var m model.UserDiaryMood
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": result.InsertedID,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToJSON(m), nil
}

func (u *UserDiaryMoodService) ToJSON(m model.UserDiaryMood) *model.UserDiaryMoodJSON {
	return &model.UserDiaryMoodJSON{
		Id:        m.Id,
		UserId:    m.UserId,
		Moods:     m.Moods,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
	}
}

func NewUserDiaryMoodService(db *mongo.Database) *UserDiaryMoodService {
	return &UserDiaryMoodService{
		db:  db,
		col: db.Collection("user_diary_mood"),
	}
}


