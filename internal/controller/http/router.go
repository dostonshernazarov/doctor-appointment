package http

import (
	"github.com/dostonshernazarov/doctor-appointment/config"
	_ "github.com/dostonshernazarov/doctor-appointment/docs" // Swagger docs.
	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/middleware"
	v1 "github.com/dostonshernazarov/doctor-appointment/internal/controller/http/v1"
	"github.com/dostonshernazarov/doctor-appointment/internal/usecase"
	"github.com/dostonshernazarov/doctor-appointment/pkg/logger"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

type Router struct {
	app         *fiber.App
	cfg         *config.Config
	l           logger.Interface
	user        usecase.UserUsecase
	doctor      usecase.DoctorUsecase
	appointment usecase.AppointmentUsecase
}

// NewRouterConfig creates a new Router configuration
func NewRouterConfig(app *fiber.App, cfg *config.Config, l logger.Interface, user usecase.UserUsecase, doctor usecase.DoctorUsecase, appointment usecase.AppointmentUsecase) *Router {
	return &Router{
		app:         app,
		cfg:         cfg,
		l:           l,
		user:        user,
		doctor:      doctor,
		appointment: appointment,
	}
}

// NewRouter -.
// Swagger spec:
// @title       Doctor appointment api
// @description API for doctor appointment
// @version     1.0
// @host        localhost:8070
// @schemes     http https
// @BasePath    /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(r *Router) {
	r.app.Use(middleware.Logger(r.l))

	// Swagger
	r.app.Get("/swagger/*", swagger.HandlerDefault)

	// Configure CORS middleware
	r.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8070, http://127.0.0.1:8070",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	// Routes
	apiV1Group := r.app.Group("/v1")
	{
		v1.NewUserRoutes(v1.HandlerV1Config{
			Config:     r.cfg,
			Logger:     r.l,
			Validation: validator.New(),
			User:       r.user,
			Router:     apiV1Group,
		})
	}

}
