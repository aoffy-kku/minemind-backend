package _interface

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"time"
)

type UserDiaryMoodServiceInterface interface {
	GetUserDiaryMoodByDate(id string, date time.Time) ([]*model.UserDiaryMoodJSON, error)
	CreateUserDiaryMood(request model.CreateUserDiaryMoodRequestJSON) (*model.UserDiaryMoodJSON, error)
	ToJSON(m model.UserDiaryMood) *model.UserDiaryMoodJSON
}
