package service

import (
	"context"
	"github.com/aoffy-kku/minemind-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleService struct {
	db  *mongo.Database
	col *mongo.Collection
}

func (r *RoleService) GetRoles() ([]*model.RoleJSON, error) {
	ctx := context.Background()
	var roles []*model.RoleJSON
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var m model.Role
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		roles = append(roles, r.ToJSON(&m))
	}
	return roles, nil
}

func (r *RoleService) ToJSON(role *model.Role) *model.RoleJSON {
	return &model.RoleJSON{
		Id:        role.Id,
		CreatedAt: role.CreatedAt,
		CreatedBy: role.CreatedBy,
	}
}

func NewRoleService(db *mongo.Database) *RoleService {
	return &RoleService{
		db:  db,
		col: db.Collection("role"),
	}
}
