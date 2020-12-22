package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type UserMoodServiceInterface interface {
	GetUserMoods(id string) ([]*model.UserMoodJSON, error)
	CreateUserMood(request model.CreateUserMoodJSON) ([]*model.UserMoodJSON, error)
	UpdateUserMood(request model.UpdateUserMoodJSON) (*model.UserMoodJSON, error)
	ToJSON(mood *model.UserMood) *model.UserMoodJSON
}
