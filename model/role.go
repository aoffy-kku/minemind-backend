package model

import "time"

type Role struct {
	Id string           `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	CreatedBy string    `bson:"created_by"`
}

type RoleJSON struct {
	Id string           `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"-"`
}