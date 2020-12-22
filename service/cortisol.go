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

type CortisolService struct {
	db *mongo.Database
	col *mongo.Collection
}

func (c *CortisolService) GetLatestCortisol(id string) (*model.CortisolJSON, error) {
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{
		"timestamp": -1,
	})
	opts.SetLimit(1)
	ctx := context.Background()
	var doc *model.CortisolJSON
	cur, err := c.col.Find(ctx, bson.M{
		"user_id": bson.M{
			"$eq": id,
		},
	}, opts)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.Cortisol
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		doc = c.ToJSON(&m)
	}
	return doc, nil
}

func (c *CortisolService) CreateCortisol(req model.CreateCortisolRequestJSON) (*model.CortisolJSON, error) {
	ctx := context.Background()
	now := time.Now()
	result, err := c.col.InsertOne(ctx, &model.Cortisol{
		Id:        primitive.NewObjectIDFromTimestamp(now),
		UserId:    req.UserId,
		Cortisol:  req.Cortisol,
		Timestamp: req.Timestamp,
		CreatedAt: now,
		CreatedBy: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	var m model.Cortisol
	if err := c.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": result.InsertedID,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return c.ToJSON(&m), nil
}

func (c *CortisolService) CreateMultipleCortisol(req model.CreateMultipleCortisol) ([]*model.CortisolJSON, error) {
	ctx := context.Background()
	now := time.Now()
	var docs []interface{}
	for _, data := range req.Data {
		docs = append(docs, &model.Cortisol{
			Id:        primitive.NewObjectIDFromTimestamp(now),
			UserId:    req.UserId,
			Cortisol:  data.Cortisol,
			Timestamp: data.Timestamp,
			CreatedAt: now,
			CreatedBy: req.UserId,
		})
	}
	result, err := c.col.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	var cortisol []*model.CortisolJSON
	cur, err := c.col.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": result.InsertedIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.Cortisol
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		cortisol = append(cortisol, c.ToJSON(&m))
	}
	return cortisol, nil
}

func (c *CortisolService) ToJSON(cortisol *model.Cortisol) *model.CortisolJSON {
	return &model.CortisolJSON{
		Id:        cortisol.Id,
		UserId:    cortisol.UserId,
		Cortisol:  cortisol.Cortisol,
		Timestamp: cortisol.Timestamp,
		CreatedAt: cortisol.CreatedAt,
		CreatedBy: cortisol.CreatedBy,
	}
}

func NewCortisolService(db *mongo.Database) *CortisolService {
	return &CortisolService{
		db:  db,
		col: db.Collection("cortisol"),
	}
}
