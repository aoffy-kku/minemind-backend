package service

import (
	"context"
	_interface "github.com/aoffy-kku/minemind-backend/interface"
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserService struct {
	db *mongo.Database
	col *mongo.Collection
	accessTokenService _interface.AccessTokenServiceInterface
}

func (u *UserService) Login(user model.UserLoginRequestJSON) (*model.AccessTokenJSON, error) {
	ctx := context.Background()
	hash, _ := utils.GeneratePassword(user.Password)
	var m model.User
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": user.Email,
		},
		"password": bson.M{
			"$eq": hash,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	result, err := u.accessTokenService.CreateToken(user.Email)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserService) Logout() error {
	panic("implement me")
}

func (u *UserService) CreateUser(user model.CreateUserRequestJSON) (*model.UserJSON, error) {
	ctx := context.Background()
	now := time.Now()
	result , err := u.col.InsertOne(ctx, model.User{
		Email:       user.Email,
		Password:    user.Password,
		DisplayName: user.DisplayName,
		WatchId:     user.WatchId,
		Roles:       []string{
			"user",
		},
		Begin:       time.Time{},
		End:         time.Time{},
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return nil, err
	}
	var m model.User
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": result.InsertedID.(string),
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToJSON(&m), nil
}

func (u *UserService) GetMe(id string) (*model.MeJSON, error) {
	ctx := context.Background()
	var m model.User
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToMeJSON(&m), nil
}

func (u *UserService) GetUserById(id string) (*model.UserJSON, error) {
	ctx := context.Background()
	var m model.User
	if err := u.col.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
	}).Decode(&m); err != nil {
		return nil, err
	}
	return u.ToJSON(&m), nil
}

func (u *UserService) GetUsers() ([]*model.UserJSON, error) {
	ctx := context.Background()
	cur, err := u.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*model.UserJSON
	for cur.Next(ctx) {
		var m model.User
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		users = append(users, u.ToJSON(&m))
	}
	return users, nil
}

func (u *UserService) ToJSON(user *model.User) *model.UserJSON {
	return &model.UserJSON{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		WatchId:     user.WatchId,
		Roles:       user.Roles,
		Begin:       user.Begin,
		End:         user.End,
		CreatedAt:   user.CreatedAt,
		CreatedBy:   user.CreatedBy,
		UpdatedAt:   user.UpdatedAt,
		UpdatedBy:   user.UpdatedBy,
	}
}

func (u *UserService) ToMeJSON(user *model.User) *model.MeJSON {
	return &model.MeJSON{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		WatchId:     user.WatchId,
		Roles:       user.Roles,
	}
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		db:  db,
		col: db.Collection("user"),
		accessTokenService: NewAccessTokenService(db),
	}
}
