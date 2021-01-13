package handler

import (
	"fmt"
	"github.com/aoffy-kku/minemind-backend/model"
	"github.com/aoffy-kku/minemind-backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

// CreateUser godoc
// @tags User
// @Summary Create user
// @Accept  json
// @Produce  json
// @Param body body model.CreateUserRequestJSON true "body"
// @Success 201 {object} model.CreateUserRequestJSON
// @Failure 400 {object} utils.HttpResponse
// @Failure 401 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users [post]
// @Security ApiKeyAuth
func (h *Handler) CreateUser(c echo.Context) error {
	var req model.CreateUserRequestJSON
	id := c.Get("id").(string)
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	req.CreatedBy = id
	hash, _ := utils.GeneratePassword(req.Password)
	req.Password = hash
	result, err := h.userService.CreateUser(req)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusUnprocessableEntity, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusCreated, result)
}

// GetMe godoc
// @tags User
// @Summary Get me
// @Accept  json
// @Produce  json
// @Success 200 {object} model.MeJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/me [get]
// @Security ApiKeyAuth
func (h *Handler) GetMe(c echo.Context) error {
	id := c.Get("id").(string)
	fmt.Println(id)
	result , err := h.userService.GetMe(id)
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetMe godoc
// @tags OldUser
// @Summary Get me
// @Accept  json
// @Produce  json
// @Success 200 {object} model.MeJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/user/me [get]
// @Security ApiKeyAuth
func (h *Handler) OldGetMe(c echo.Context) error {
	id := c.Get("id").(string)
	fmt.Println(id)
	result , err := h.userService.GetMe(id)
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUserById godoc
// @tags User
// @Summary Get user by id
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} model.UserJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	result , err := h.userService.GetUserById(id)
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, result)
}

// GetUsers godoc
// @tags User
// @Summary Get users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.UserJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsers(c echo.Context) error  {
	results, err := h.userService.GetUsers()
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, results)
}

// Login godoc
// @tags User
// @Summary Login
// @Accept  json
// @Produce  json
// @Param body body model.UserLoginRequestJSON true "body"
// @Success 200 {array} model.AccessTokenJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/login [post]
func (h *Handler) Login(c echo.Context) error  {
	var req model.UserLoginRequestJSON
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	results, err := h.userService.Login(req)
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, results)
}

// Login godoc
// @tags Public
// @Summary Login
// @Accept  json
// @Produce  json
// @Param body body model.UserLoginRequestJSON true "body"
// @Success 200 {array} model.AccessTokenJSON
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/login [post]
func (h *Handler) OldLogin(c echo.Context) error  {
	var req model.UserLoginRequestJSON
	if err := c.Bind(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	results, err := h.userService.Login(req)
	if err !=nil {
		if err == mongo.ErrNoDocuments {
			return utils.EchoHttpResponse(c, http.StatusNotFound, utils.HttpResponse{})
		}
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	type OldAccessToken struct {
		AccessToken string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	return utils.EchoHttpResponse(c, http.StatusOK, &OldAccessToken{
		AccessToken: results.Pair,
		RefreshToken: results.Id,
	})
}

// UpdateBirthDate godoc
// @tags User
// @Summary Update birth date
// @Accept  json
// @Produce  json
// @Param birthdate path string true "birthdate"
// @Success 200 {object} utils.HttpResponse
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/users/birthdate/{birthdate} [patch]
// @Security ApiKeyAuth
func (h *Handler) UpdateBirthDate(c echo.Context) error  {
	date := c.Param("birthdate")
	id := c.Get("id").(string)
	ts, err := time.Parse("2006-01-02T03-04-05Z", date)
	if err != nil {
		return utils.EchoHttpResponse(c, http.StatusBadRequest, utils.HttpResponse{Message: err.Error()})
	}
	if err := h.userService.UpdateBirthDate(id, ts); err !=nil {
		return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
	}
	return utils.EchoHttpResponse(c, http.StatusOK, utils.HttpResponse{})
}