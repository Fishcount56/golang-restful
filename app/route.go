package app

import (
	"database/sql"
	"restful-api/controller"
	"restful-api/providers"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB) *httprouter.Router {
	router := httprouter.New()

	validate := validator.New()

	servicesProviders := providers.NewServiceProvider(db, validate)
	controllerProviders := providers.NewControllerProvider(servicesProviders)

	router.GET("/health-check", controller.HealthController)
	router.POST("/product", controllerProviders.ProductController.CreateProduct)

	return router
}
