package handler

import (
	"github.com/aoffy-kku/minemind-backend/router/middleware"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	auth := middleware.JWT([]byte("Minemind2019"), h.db)

	user := v1.Group("/users")
	user.POST("", h.CreateUser)
	user.GET("/me", h.GetMe, auth)
	user.GET("/:id", h.GetUserById, auth)
	user.GET("", h.GetUsers, auth)

	role := v1.Group("/roles")
	role.GET("", h.GetRoles)

}