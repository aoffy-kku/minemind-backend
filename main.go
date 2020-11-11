package main

import (
	db2 "github.com/aoffy-kku/minemind-backend/db"
	_ "github.com/aoffy-kku/minemind-backend/docs"
	"github.com/aoffy-kku/minemind-backend/handler"
	"github.com/aoffy-kku/minemind-backend/router"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description MineMind API
// @title MineMind API

// @BasePath /

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := r.Group("/v1")
	db := db2.New()
	h := handler.NewHandler(db)
	h.Register(v1)
	log.Fatal(r.Start("127.0.0.1:1321"))
}
