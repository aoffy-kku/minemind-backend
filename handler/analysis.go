package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// CreateAnalysis godoc
// @tags Analysis
// @Summary Create analysis
// @Accept  json
// @Produce  json
// @Param body body model.CreateAnalysisRequestJSON true "body"
// @Success 201 {object} utils.HttpResponse
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/analysis [post]
// @Security ApiKeyAuth
func (h *Handler) CreateAnalysis(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateAnalysisRequestJSON
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
	if err := h.analysisService.CreateAnalysis(req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusCreated, utils.HttpResponse{})
}

// GetAnalysisByDate godoc
// @tags Analysis
// @Summary Get analysis by date
// @Accept  json
// @Produce  json
// @Param date query string false "date"
// @Success 200 {array} model.AnalysisJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/analysis [get]
// @Security ApiKeyAuth
func (h *Handler) GetAnalysisByDate(c echo.Context) error {
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
	result, err := h.analysisService.GetAnalysisByDate(id, time.Unix(dt, 0))
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetAnalysisByMode godoc
// @tags Analysis
// @Summary Get analysis by mode
// @Accept  json
// @Produce  json
// @Param mode query string false "mode"
// @Success 200 {array} model.AnalysisJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/analysis/mode/{mode} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAnalysisByMode(c echo.Context) error {
	id := c.Get("id").(string)
	mode := c.Param("mode")
	m, err := strconv.ParseInt(mode, 10, 64)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	result, err := h.analysisService.GetAnalysisByMode(id, m)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

