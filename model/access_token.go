package model

import "time"

type AccessToken struct {
	Id string           `bson:"_id"`
	UserId string       `bson:"user_id"`
	Pair string         `bson:"pair"`
	CreatedAt time.Time `bson:"created_at"`
	CreatedBy string    `bson:"created_by"`
}
type UpdateAccessTokenJSON struct {
	Id string   `json:"id"`
	Pair string `json:"-"`
}
type AccessTokenJSON struct {
	Id string           `json:"id"`
	UserId string       `json:"userId"`
	Pair string         `json:"pair"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}