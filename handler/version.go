package handler

import (
    "context"
    "github.com/aoffy-kku/minemind-backend/utils"
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson"
    "net/http"
)

type Version struct {
    Code string `bson:"code" json:"code"`
    Url string  `bson:"url" json:"url"`
}
// GetVersion godoc
// @tags Version
// @Summary Get version
// @Accept  json
// @Produce  json
// @Success 200 {object} Version
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/version [get]
func (h *Handler) GetVersion(c echo.Context) error {
    var m Version
    if err := h.db.Collection("version").FindOne(context.Background(), bson.M{
        "_id": bson.M{
            "$eq": "minemind",
        },
    }).Decode(&m); err != nil {
        return utils.EchoHttpResponse(c, http.StatusInternalServerError, utils.HttpResponse{Message: err.Error()})
    }
    return utils.EchoHttpResponse(c, http.StatusOK, m)
}