package _interface

import "github.com/aoffy-kku/minemind-backend/model"

type AdminServiceInterface interface {
    GetUsersDiary() ([]*model.UserDiaryJSON, error)
    GetUsersEvaluation() ([]*model.UserEvaluationJSON, error)
    GetUsersCortisol() ([]*model.CortisolJSON, error)
    GetUserDiary(id string) ([]*model.UserDiaryJSON, error)
    GetUserEvaluation(id string) ([]*model.UserEvaluationJSON, error)
    GetUserCortisol(id string) ([]*model.CortisolJSON, error)
    GetUsers() ([]*model.UserJSON, error)
    UpdateUser(request model.UpdateUserRequestJSON) (*model.UserJSON, error)
}