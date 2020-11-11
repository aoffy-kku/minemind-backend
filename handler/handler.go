package handler

import (
	_interface "github.com/aoffy-kku/minemind-backend/interface"
	"github.com/aoffy-kku/minemind-backend/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	db *mongo.Database
	userService _interface.UserServiceInterface
	roleService _interface.RoleServiceInterface
	accessTokenService _interface.AccessTokenServiceInterface
}

func NewHandler(db *mongo.Database) *Handler {
	var userService _interface.UserServiceInterface = service.NewUserService(db)
	var roleService _interface.RoleServiceInterface = service.NewRoleService(db)
	var accessTokenService _interface.AccessTokenServiceInterface = service.NewAccessTokenService(db)
	return &Handler{
		db: db,
		userService: userService,
		roleService: roleService,
		accessTokenService: accessTokenService,
	}
}
