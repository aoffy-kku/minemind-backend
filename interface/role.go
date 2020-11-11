package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type RoleServiceInterface interface {
	GetRoles() ([]*model.RoleJSON, error)
	ToJSON(role *model.Role) *model.RoleJSON
}