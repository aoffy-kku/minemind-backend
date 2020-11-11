package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type AccessTokenServiceInterface interface {
	GetAccessToken(req model.UpdateAccessTokenJSON) (*model.AccessTokenJSON, error)
	CreateToken(id string) (*model.AccessTokenJSON, error)
	ToJSON(token *model.AccessToken) *model.AccessTokenJSON
}
