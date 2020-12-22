package _interface

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"time"
)

type UserDiaryServiceInterface interface {
	GetUserDiaryByDate(id string, date time.Time) ([]*model.UserDiaryJSON, error)
	CreateUserDiary(request model.CreateUserDiaryRequestJSON) (*model.UserDiaryJSON, error)
	ToJSON(m model.UserDiary) *model.UserDiaryJSON
}
