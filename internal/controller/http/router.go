package http

import (
	"github.com/dostonshernazarov/doctor-appointment/config"
	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/middleware"
	"github.com/dostonshernazarov/doctor-appointment/pkg/logger"
	"github.com/gofiber/fiber"
)

// NewRouter -.
// Swagger spec:
// @title       Doctor appointment api
// @description API for doctor appointment
// @version     1.0
// @host        localhost:8070
// @BasePath    /v1
func NewRouter(app *fiber.App, cfg *config.Config, l logger.Interface, uc usecase.UseCase) {
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recover(l))

	// Prometheus metrics
	if cfg.Metrics.Enabled {
		prometheus := fiberprometheus.New("doctor-appointment")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// Routes
	apiV1Group := app.Group("/v1")
	{
		v1.NewUserRoutes(apiV1Group, uc.User, l)
	}

}
