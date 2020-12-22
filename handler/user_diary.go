package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// CreateUserDiary godoc
// @tags UserDiary
// @Summary Create user diary
// @Accept  json
// @Produce  json
// @Param body body model.CreateUserDiaryRequestJSON true "body"
// @Success 200 {array} model.UserDiaryJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/diary [post]
// @Security ApiKeyAuth
func (h *Handler) CreateUserDiary(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateUserDiaryRequestJSON
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
	result, err := h.userDiaryService.CreateUserDiary(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserDiaryByDate godoc
// @tags UserDiary
// @Summary Get user diary by date
// @Accept  json
// @Produce  json
// @Param date query string false "date"
// @Success 200 {array} model.UserDiaryJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/diary [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserDiaryByDate(c echo.Context) error {
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
	result, err := h.userDiaryService.GetUserDiaryByDate(id, time.Unix(dt, 0))
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}
