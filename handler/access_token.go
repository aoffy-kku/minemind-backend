package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAccessToken godoc
// @tags AccessToken
// @Summary Get access token
// @Accept  json
// @Produce  json
// @Param body body model.UpdateAccessTokenJSON true "body"
// @Success 200 {array} model.AccessTokenJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/access-token/renew [post]
func (h *Handler) GetAccessToken(c echo.Context) error {
	var req model.UpdateAccessTokenJSON
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	result, err := h.accessTokenService.GetAccessToken(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}
