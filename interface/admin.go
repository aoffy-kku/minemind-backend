package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type AdminServiceInterface interface {
    GetUsersDiary() ([]*model.UserDiaryJSON, error)
    GetUsersEvaluation() ([]*model.UserEvaluationJSON, error)
    GetUsersCortisol() ([]*model.CortisolJSON, error)
    UpdateUser(request model.UpdateUserRequestJSON) (*model.UserJSON, error)
}