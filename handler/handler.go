package handler

import (
	_interface "github.com/aoffy-kku/minemind-backend/interface"
	"github.com/aoffy-kku/minemind-backend/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	db *mongo.Database
	userService _interface.UserServiceInterface
	roleService _interface.RoleServiceInterface
	accessTokenService _interface.AccessTokenServiceInterface
	cortisolService _interface.CortisolServiceInterface
	analysisService _interface.AnalysisServiceInterface
	userMoodService _interface.UserMoodServiceInterface
	userDiaryMoodService _interface.UserDiaryMoodServiceInterface
	userDiaryService _interface.UserDiaryServiceInterface
	evaluationService _interface.EvaluationServiceInterface
	userEvaluationService _interface.UserEvaluationServiceInterface
}

func NewHandler(db *mongo.Database) *Handler {
	var userService _interface.UserServiceInterface = service.NewUserService(db)
	var roleService _interface.RoleServiceInterface = service.NewRoleService(db)
	var accessTokenService _interface.AccessTokenServiceInterface = service.NewAccessTokenService(db)
	var cortisolService _interface.CortisolServiceInterface = service.NewCortisolService(db)
	var analysisService _interface.AnalysisServiceInterface = service.NewAnalysisService(db)
	var userMoodService _interface.UserMoodServiceInterface = service.NewUserMoodService(db)
	var userDiaryMoodService _interface.UserDiaryMoodServiceInterface = service.NewUserDiaryMoodService(db)
	var userDiaryService _interface.UserDiaryServiceInterface = service.NewUserDiaryService(db)
	var evaluationService _interface.EvaluationServiceInterface = service.NewEvaluationService(db)
	var userEvaluationService _interface.UserEvaluationServiceInterface = service.NewUserEvaluationService(db)

	return &Handler{
		db: db,
		userService: userService,
		roleService: roleService,
		accessTokenService: accessTokenService,
		cortisolService: cortisolService,
		analysisService: analysisService,
		userMoodService: userMoodService,
		userDiaryMoodService: userDiaryMoodService,
		userDiaryService: userDiaryService,
		evaluationService: evaluationService,
		userEvaluationService: userEvaluationService,
	}
}
