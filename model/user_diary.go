package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserDiary struct {
	Id primitive.ObjectID      `bson:"_id"`
	UserId string              `bson:"user_id"`
	Moods []primitive.ObjectID `bson:"moods"`
	Content string             `bson:"content"`
	CreatedAt time.Time        `bson:"created_at"`
	CreatedBy string           `bson:"created_by"`
}

type UserDiaryJSON struct {
	Id primitive.ObjectID      `json:"id"`
	UserId string              `json:"userId"`
	Moods []primitive.ObjectID `json:"moods"`
	Content string             `json:"content"`
	CreatedAt time.Time        `json:"createdAt"`
	CreatedBy string           `json:"createdBy"`
}

type CreateUserDiaryRequestJSON struct {
	UserId string   `json:"-"`
	Moods []primitive.ObjectID  `json:"moods"`
	Content string  `json:"content"`
}

type UpdateUserDiaryRequestJSON struct {
	UserId string   `json:"-"`
	Moods []primitive.ObjectID  `json:"moods"`
	Content string  `json:"content"`
}
