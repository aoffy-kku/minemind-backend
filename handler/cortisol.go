package handler

import (
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateCortisol godoc
// @tags Cortisol
// @Summary Create cortisol
// @Accept  json
// @Produce  json
// @Param body body model.CreateCortisolRequestJSON true "body"
// @Success 201 {object} model.CortisolJSON
// @Failure 400 {object} utils.HttpResponse
// @Failure 401 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/cortisol [post]
func (h *Handler) CreateCortisol(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateCortisolRequestJSON
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	req.UserId = id
	result, err := h.cortisolService.CreateCortisol(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusCreated, result)
}

// CreateMultipleCortisol godoc
// @tags Cortisol
// @Summary Create multiple cortisol
// @Accept  json
// @Produce  json
// @Param body body model.CreateMultipleCortisol true "body"
// @Success 201 {array} model.CortisolJSON
// @Failure 400 {object} utils.HttpResponse
// @Failure 401 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/cortisol/backup [post]
func (h *Handler) CreateMultipleCortisol(c echo.Context) error {
	id := c.Get("id").(string)
	var req model.CreateMultipleCortisol
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	req.UserId = id
	result, err := h.cortisolService.CreateMultipleCortisol(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
			Message: err.Error(),
		})
	}
	return utils.EchoHttpResponse(c, http.StatusCreated, result)
}
