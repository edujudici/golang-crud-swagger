package main

import (
	"net/http"

	"github.com/edujudici/golang-crud-swagger/config"
	"github.com/edujudici/golang-crud-swagger/controller"
	_ "github.com/edujudici/golang-crud-swagger/docs"
	"github.com/edujudici/golang-crud-swagger/helper"
	"github.com/edujudici/golang-crud-swagger/model"
	"github.com/edujudici/golang-crud-swagger/repository"
	"github.com/edujudici/golang-crud-swagger/router"
	"github.com/edujudici/golang-crud-swagger/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
