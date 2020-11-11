package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserDiary struct {
	Id primitive.ObjectID
	UserId string
	Title string
	IsSpecial bool
	Moods []primitive.ObjectID
	Content string
	Timestamp int64
	CreatedAt time.Time
	CreatedBy string
}

type CreateUserDiaryRequestJSON struct {
	Title string
	IsSpecial bool
	Moods []string
	Content string
}


type UpdateUserDiaryRequestJSON struct {
	Title string
	IsSpecial bool
	Moods []string
	Content string
}
