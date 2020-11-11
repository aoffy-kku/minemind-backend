package handler

import (
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// GetRoles godoc
// @tags Roles
// @Summary Get roles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.RoleJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/roles [get]
func (h *Handler) GetRoles(c echo.Context) error  {
	results, err := h.roleService.GetRoles()
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, results)
}