package _interface

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserDiaryServiceInterface interface {
	GetUserDiaryByDate(id string, date time.Time) ([]*model.UserDiaryJSON, error)
	CreateUserDiary(request model.CreateUserDiaryRequestJSON) (*model.UserDiaryJSON, error)
	DeleteUserDiary(id primitive.ObjectID, uid string) error
	ToJSON(m model.UserDiary) *model.UserDiaryJSON
}
