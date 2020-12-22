package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateUserEvaluation godoc
// @tags UserEvaluation
// @Summary Create user evaluation
// @Accept  json
// @Produce  json
// @Param body body model.CreateUserEvaluationRequestJSON true "body"
// @Success 200 {object} model.UserEvaluationJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/evaluation [post]
// @Security ApiKeyAuth
func (h *Handler) CreateUserEvaluation(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateUserEvaluationRequestJSON
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	if err := c.Validate(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	req.UserId = id
	result, err := h.userEvaluationService.CreateUserEvaluation(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetLatestUserEvaluation godoc
// @tags UserEvaluation
// @Summary Get latest user evaluation
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserEvaluationJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/evaluation [get]
// @Security ApiKeyAuth
func (h *Handler) GetLatestUserEvaluation(c echo.Context) error {
	id := c.Get("id").(string)
	result, err := h.userEvaluationService.GetLatestUserEvaluation(id)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}
