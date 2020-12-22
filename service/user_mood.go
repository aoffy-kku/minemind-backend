package service

import (
	"context"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserMoodService struct {
	db *mongo.Database
	col *mongo.Collection
}

func (u *UserMoodService) GetUserMoods(id string) ([]*model.UserMoodJSON, error) {
	ctx := context.Background()
	var moods []*model.UserMoodJSON
	cur, err := u.col.Find(ctx, bson.M{
		"user_id": bson.M{
			"$eq": id,
		},
	})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.UserMood
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		moods = append(moods, u.ToJSON(&m))
	}
	return moods, nil
}

func (u *UserMoodService) CreateUserMood(request model.CreateUserMoodJSON) ([]*model.UserMoodJSON, error) {
	ctx := context.Background()
	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)
	for _, mood := range request.Moods {
		_, err := u.col.UpdateOne(ctx, bson.M{
			"name": bson.M{
				"$eq": mood,
			},
		}, bson.M{
			"$set": bson.M{
				"name": mood,
				"active": true,
				"created_at": time.Now(),
				"created_by": "system",
			},
		}, opts)
		if err != nil {
			return nil, err
		}
	}
	var moods []*model.UserMoodJSON
	cur, err := u.col.Find(ctx, bson.M{
		"name": bson.M{
			"$in": request.Moods,
		},
	})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.UserMood
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		moods = append(moods, u.ToJSON(&m))
	}
	return moods, nil
}

func (u *UserMoodService) UpdateUserMood(request model.UpdateUserMoodJSON) (*model.UserMoodJSON, error) {
	ctx := context.Background()
	var m model.UserMood
	if err := u.col.FindOneAndUpdate(ctx, bson.M{
		"_id": bson.M{
			"$eq": request.Id,
		},
	}, bson.M{
		"$set": bson.M{
			"name": request.Name,
			"active": request.Active,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToJSON(&m), nil
}

func (u *UserMoodService) ToJSON(mood *model.UserMood) *model.UserMoodJSON {
	return &model.UserMoodJSON{
		Id:        mood.Id,
		Name:      mood.Name,
		UserId:    mood.UserId,
		Active:    mood.Active,
		CreatedAt: mood.CreatedAt,
		CreatedBy: mood.CreatedBy,
	}
}

func NewUserMoodService(db *mongo.Database) *UserMoodService {
	return &UserMoodService{
		db:  db,
		col: db.Collection("user_mood"),
	}
}