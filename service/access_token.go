package service

import (
	"context"
	"fmt"
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AccessTokenService struct {
	db *mongo.Database
	col *mongo.Collection
}

func (a *AccessTokenService) GetAccessToken(req model.UpdateAccessTokenJSON) (*model.AccessTokenJSON, error) {
	ctx := context.Background()
	count, err := a.col.CountDocuments(ctx, bson.M{
		"_id": bson.M{
			"$eq": req.Id,
		},
		"pair": bson.M{
			"$eq": req.Pair,
		},
	})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("invalid token")
	}
	accessToken := utils.GenerateAccessToken(req.Id, 1000 * 60 * 60)
	_, err = a.col.UpdateOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": req.Id,
		},
	}, bson.M{
		"$set": bson.M{
			"pair": accessToken,
		},
	})
	if err != nil {
		return nil, err
	}
	var m model.AccessToken
	if err := a.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": req.Id,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return a.ToJSON(&m), nil
}

func (a *AccessTokenService) CreateToken(id string) (*model.AccessTokenJSON, error) {
	ctx := context.Background()
	refreshToken := utils.GenerateAccessToken(id, 1000 * 60 * 60 * 24 * 365)
	accessToken := utils.GenerateAccessToken(id, 1000 * 60 * 60)
	result, err := a.col.InsertOne(ctx, model.AccessToken{
		Id:        refreshToken,
		UserId:    id,
		Pair:      accessToken,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	var m model.AccessToken
	if err := a.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": result.InsertedID.(string),
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return a.ToJSON(&m), nil
}

func (a *AccessTokenService) ToJSON(token *model.AccessToken) *model.AccessTokenJSON {
	return &model.AccessTokenJSON{
		Id:        token.Id,
		UserId:    token.UserId,
		Pair:      token.Pair,
		CreatedAt: token.CreatedAt,
		CreatedBy: token.CreatedBy,
	}
}

func NewAccessTokenService(db *mongo.Database) *AccessTokenService {
	return &AccessTokenService{
		db:  db,
		col: db.Collection("access_token"),
	}
}
