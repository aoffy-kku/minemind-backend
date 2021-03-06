package handler

import (
	"github.com/aoffy-kku/minemind-backend/router/middleware"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	auth := middleware.JWT([]byte("Minemind2019"), h.db)

	public := v1.Group("")
	public.POST("/login", h.OldLogin)
	public.POST("/token", h.OldGetAccessToken)

	oldUser := v1.Group("/user")
	oldUser.GET("/me", h.OldGetMe, auth)

	user := v1.Group("/users")
	user.POST("", h.CreateUser, auth, middleware.Roles(h.db, middleware.Admin))
	user.GET("/me", h.GetMe, auth)
	user.PATCH("/birthdate/:birthdate", h.UpdateBirthDate, auth)
	user.GET("/:id", h.GetUserById, auth, middleware.Roles(h.db, middleware.Admin))
	user.GET("", h.GetUsers, auth, middleware.Roles(h.db, middleware.Admin))
	user.POST("/login", h.Login)
	// user mood
	user.GET("/mood", h.GetUserMoods, auth)
	user.POST("/mood", h.CreateUserMood, auth)
	user.PATCH("/mood/:mood_id", h.UpdateUserMood, auth)
	// user diary mood
	//user.POST("/diary/mood", h.CreateUserDiaryMood, auth)
	//user.GET("/diary/mood", h.GetUserDiaryMoodByDate, auth)
	// user diary
	user.POST("/diary", h.CreateUserDiary, auth)
	user.GET("/diary", h.GetUserDiaryByDate, auth)
	user.DELETE("/diary/:id", h.DeleteUserDiary, auth)
	// user analysis
	user.POST("/analysis", h.CreateAnalysis, auth)
	user.GET("/analysis", h.GetAnalysisByDate, auth)
	user.DELETE("/analysis/:id", h.DeleteAnalysis, auth)
	user.GET("/analysis/mode/:mode", h.GetAnalysisByMode, auth)

	// user evaluation
	user.POST("/evaluation", h.CreateUserEvaluation, auth)
	user.GET("/evaluation", h.GetLatestUserEvaluation, auth)

	// user cortisol
	user.GET("/cortisol/latest", h.GetLatestCortisol, auth)

	role := v1.Group("/roles")
	role.GET("", h.GetRoles)

	evaluation := v1.Group("/evaluations")
	evaluation.GET("", h.GetEvaluations)

	cortisol := v1.Group("/cortisol", auth)
	cortisol.POST("", h.CreateCortisol)
	cortisol.POST("/backup", h.CreateMultipleCortisol)

	accessToken := v1.Group("/access-token")
	accessToken.POST("/renew", h.GetAccessToken)

	measurement := v1.Group("/measurement", auth)
	measurement.GET("", h.GetMeasurement)

	version := v1.Group("/version")
	version.GET("", h.GetVersion)

	admin := v1.Group("/admin", auth, middleware.Roles(h.db, middleware.Admin))
	admin.GET("/users/diary", h.GetUsersDiary)
	admin.GET("/users/evaluation", h.GetUsersEvaluations)
	admin.GET("/users/cortisol", h.GetUsersCortisol)
	admin.GET("/users/analysis", h.GetUsersAnalysis)
	admin.GET("/users/:id/diary", h.GetUserDiary)
	admin.GET("/users/:id/evaluation", h.GetUserEvaluations)
	admin.GET("/users/:id/cortisol", h.GetUserCortisol)
	admin.GET("/users/:id/analysis", h.GetUserAnalysis)
	admin.GET("/users/:id/mood", h.GetUserMood)
	admin.GET("/users", h.GetUsers)
	admin.PATCH("/users/:id", h.UpdateUser)
}