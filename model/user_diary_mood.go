package model

import (
"go.mongodb.org/mongo-driver/bson/primitive"
"time"
)

type UserDiaryMood struct {
	Id primitive.ObjectID      `bson:"_id"`
	UserId string              `bson:"user_id"`
	Moods []string `bson:"moods"`
	CreatedAt time.Time        `bson:"created_at"`
	CreatedBy string           `bson:"created_by"`
}

type UserDiaryMoodJSON struct {
	Id primitive.ObjectID      `json:"id"`
	UserId string              `json:"userId"`
	Moods []string `json:"moods"`
	CreatedAt time.Time        `json:"createdAt"`
	CreatedBy string           `json:"createdBy"`
}

type CreateUserDiaryMoodRequestJSON struct {
	UserId string   `json:"-"`
	Moods []string  `json:"moods"`
}

type UpdateUserDiaryMoodRequestJSON struct {
	UserId string   `json:"-"`
	Moods []string  `json:"moods"`
}
