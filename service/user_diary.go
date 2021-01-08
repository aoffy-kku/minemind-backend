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

type UserDiaryService struct {
	db *mongo.Database
	col *mongo.Collection
}

func (u *UserDiaryService) DeleteUserDiary(id primitive.ObjectID, uid string) error {
	ctx := context.Background()
	_, err := u.col.DeleteOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
		"user_id": bson.M{
			"$eq": uid,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDiaryService) GetUserDiaryByDate(id string, date time.Time) ([]*model.UserDiaryJSON, error) {
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
	var docs []*model.UserDiaryJSON
	for cur.Next(ctx) {
		var m model.UserDiary
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		docs = append(docs, u.ToJSON(m))
	}
	return docs, nil
}

func (u *UserDiaryService) CreateUserDiary(request model.CreateUserDiaryRequestJSON) (*model.UserDiaryJSON, error) {
	ctx := context.Background()
	now := time.Now()
	result, err := u.col.InsertOne(ctx, &model.UserDiary{
		Id:        primitive.NewObjectIDFromTimestamp(now),
		UserId:    request.UserId,
		Moods:     request.Moods,
		Content: request.Content,
		CreatedAt: now,
		CreatedBy: request.UserId,
	})
	if err != nil {
		return nil, err
	}
	var m model.UserDiary
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": result.InsertedID,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToJSON(m), nil
}

func (u *UserDiaryService) ToJSON(m model.UserDiary) *model.UserDiaryJSON {
	return &model.UserDiaryJSON{
		Id:        m.Id,
		UserId:    m.UserId,
		Moods:     m.Moods,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
	}
}

func NewUserDiaryService(db *mongo.Database) *UserDiaryService {
	return &UserDiaryService{
		db:  db,
		col: db.Collection("user_diary"),
	}
}