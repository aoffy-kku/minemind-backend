package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// CreateUserDiaryMood godoc
// @tags UserDiaryMood
// @Summary Create user diary mood
// @Accept  json
// @Produce  json
// @Param body body model.CreateUserDiaryMoodRequestJSON true "body"
// @Success 200 {array} model.UserDiaryMoodJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/diary/mood [post]
// @Security ApiKeyAuth
func (h *Handler) CreateUserDiaryMood(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateUserDiaryMoodRequestJSON
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
	result, err := h.userDiaryMoodService.CreateUserDiaryMood(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserDiaryMoodByDate godoc
// @tags UserDiaryMood
// @Summary Get user diary mood by date
// @Accept  json
// @Produce  json
// @Param date query string false "date"
// @Success 200 {array} model.UserDiaryMoodJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/diary/mood [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserDiaryMoodByDate(c echo.Context) error {
	id := c.Get("id").(string)
	date := c.QueryParam("date")
	var dt int64
	if len(date) == 0 {
		dt = time.Now().Unix()
	} else	if len(date) != 10 {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: "invalid date",
		})
	} else {
		unix, err := strconv.ParseInt(date, 10, 64)
		if err != nil {
			return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
				Message: err.Error(),
			})
		}
		dt = unix
	}
	result, err := h.userDiaryMoodService.GetUserDiaryMoodByDate(id, time.Unix(dt, 0))
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}
