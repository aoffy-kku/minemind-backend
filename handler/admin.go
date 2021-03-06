package handler

import (
    "github.com/aoffy-kku/minemind-backend/model"
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

// GetUsersAnalysis godoc
// @tags Admin
// @Summary Get users analysis
// @Accept  json
// @Produce  json
// @Success 200 {array} model.AnalysisJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/analysis [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsersAnalysis(c echo.Context) error {
    result, err := h.adminService.GetUsersAnalysis()
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserCortisol godoc
// @tags Admin
// @Summary Get user cortisol
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {array} model.CortisolJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/{id}/cortisol [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserCortisol(c echo.Context) error {
    id := c.Param("id")
    result, err := h.adminService.GetUserCortisol(id)
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserDiary godoc
// @tags Admin
// @Summary Get user diary
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {array} model.UserDiaryJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/{id}/diary [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserDiary(c echo.Context) error {
    id := c.Param("id")
    result, err := h.adminService.GetUserDiary(id)
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserEvaluations godoc
// @tags Admin
// @Summary Get user evaluations
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {array} model.UserEvaluationJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/{id}/evaluation [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserEvaluations(c echo.Context) error {
    id := c.Param("id")
    result, err := h.adminService.GetUserEvaluation(id)
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserAnalysis godoc
// @tags Admin
// @Summary Get users analysis
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {array} model.AnalysisJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/{id}/analysis [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserAnalysis(c echo.Context) error {
    id := c.Param("id")
    result, err := h.adminService.GetUserAnalysis(id)
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUsers godoc
// @tags Admin
// @Summary Get users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsers(c echo.Context) error {
    result, err := h.adminService.GetUsers()
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// UpdateUser godoc
// @tags Admin
// @Summary Update user
// @Accept  json
// @Produce  json
// @Param body body model.UpdateUserRequestJSON true "body"
// @Success 200 {object} model.UserJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/{id} [patch]
// @Security ApiKeyAuth
func (h *Handler) UpdateUser(c echo.Context) error {
    id := c.Param("id")
    if _, err := h.userService.GetUserById(id); err != nil {
        return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    var body model.UpdateUserRequestJSON
    if err := c.Bind(&body); err != nil {
        return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    if err := c.Validate(&body); err != nil {
        return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    body.Email = id
    result, err := h.adminService.UpdateUser(body)
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserMood godoc
// @tags Admin
// @Summary Get user mood
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {array} model.UserMoodJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/admin/users/{id}/mood [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserMood(c echo.Context) error {
    id := c.Param("id")
    result, err := h.userMoodService.GetUserMoods(id)
    if err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{
            Message: err.Error(),
        })
    }
    return utils.EchoHttpResponse(c, http.StatusOK, result)
}