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
	var token model.AccessToken
	if err := a.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": req.Id,
		},
		"pair": bson.M{
			"$eq": req.Pair,
		},
	}).Decode(&token); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}
	accessToken := utils.GenerateAccessToken(token.UserId, time.Now().Add(time.Hour * 24 * 30).Unix())
	_, err := a.col.UpdateOne(ctx, bson.M{
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
	refreshToken := utils.GenerateAccessToken(id, time.Now().Add(time.Hour * 24 * 30 * 365).Unix())
	accessToken := utils.GenerateAccessToken(id, time.Now().Add(time.Hour * 24 * 30).Unix())
	result, err := a.col.InsertOne(ctx, model.AccessToken{
		Id:        refreshToken,
		UserId:    id,
		Pair:      accessToken,
		CreatedAt: time.Now(),
		CreatedBy: "system",
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
	if _, err := a.db.Collection("user").UpdateOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
		"begin": bson.M{
			"$eq": time.Time{},
		},
	}, bson.M{
		"$set": bson.M{
			"begin": time.Now(),
		},
	}); err != nil {
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
