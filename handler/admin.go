package handler

import (
    "github.com/aoffy-kku/minemind-backend/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

// GetUsersDiary godoc
// @tags Admin
// @Summary Get users diary
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserDiaryJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/diary [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsersDiary(c echo.Context) error {
    result, err := h.adminService.GetUsersDiary()
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUsersEvaluations godoc
// @tags Admin
// @Summary Get users evaluations
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserEvaluationJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/evaluation [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsersEvaluations(c echo.Context) error {
    result, err := h.adminService.GetUsersEvaluation()
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUsersCortisol godoc
// @tags Admin
// @Summary Get users cortisol
// @Accept  json
// @Produce  json
// @Success 200 {array} model.CortisolJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/cortisol [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsersCortisol(c echo.Context) error {
    result, err := h.adminService.GetUsersCortisol()
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}