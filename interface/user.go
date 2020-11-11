package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type UserServiceInterface interface {
	CreateUser(user model.CreateUserRequestJSON) (*model.UserJSON, error)
	GetMe(id string) (*model.MeJSON, error)
	GetUserById(id string) (*model.UserJSON, error)
	GetUsers() ([]*model.UserJSON, error)
	Login(user model.UserLoginRequestJSON) (*model.AccessTokenJSON, error)
	Logout() error
	ToJSON(user *model.User) *model.UserJSON
	ToMeJSON(user *model.User) *model.MeJSON
}
