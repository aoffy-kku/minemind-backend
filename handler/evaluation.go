package handler

import (
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// GetEvaluations godoc
// @tags Evaluation
// @Summary Get evaluations
// @Accept  json
// @Produce  json
// @Success 200 {array} model.EvaluationJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/evaluations [get]
func (h *Handler) GetEvaluations(c echo.Context) error  {
	results, err := h.evaluationService.GetEvaluations()
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, results)
}
