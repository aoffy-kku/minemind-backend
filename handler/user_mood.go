package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetUserMoods godoc
// @tags UserMood
// @Summary Get user moods
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserMoodJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/mood [get]
// @Security ApiKeyAuth
func (h Handler) GetUserMoods(c echo.Context) error {
	id := c.Get("id").(string)
	result, err := h.userMoodService.GetUserMoods(id)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// CreateUserMood godoc
// @tags UserMood
// @Summary Create user moods
// @Accept  json
// @Produce  json
// @Param body body model.CreateUserMoodJSON true "body"
// @Success 200 {array} model.UserMoodJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/mood [post]
// @Security ApiKeyAuth
func (h Handler) CreateUserMood(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateUserMoodJSON
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
	result, err := h.userMoodService.CreateUserMood(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// UpdateUserMood godoc
// @tags UserMood
// @Summary Update user moods
// @Accept  json
// @Produce  json
// @Param mood_id path string true "mood_id"
// @Param body body model.UpdateUserMoodJSON true "body"
// @Success 200 {object} model.UserMoodJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/mood/{mood_id} [patch]
// @Security ApiKeyAuth
func (h Handler) UpdateUserMood(c echo.Context) error {
	mid := c.Param("mood_id")
	oid, err := primitive.ObjectIDFromHex(mid)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	var req model.UpdateUserMoodJSON
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
	req.Id = oid
	result, err := h.userMoodService.UpdateUserMood(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}