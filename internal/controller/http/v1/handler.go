package v1

import (
	"time"

	"github.com/dostonshernazarov/doctor-appointment/config"
	"github.com/dostonshernazarov/doctor-appointment/internal/usecase"
	"github.com/dostonshernazarov/doctor-appointment/pkg/logger"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type HandlerV1 struct {
	Config         *config.Config
	Logger         logger.Interface
	Validation     *validator.Validate
	ContextTimeout time.Duration
	User           usecase.UserUsecase
	Doctor         usecase.DoctorUsecase
	Appointment    usecase.AppointmentUsecase
	Router         fiber.Router
}

type HandlerV1Config struct {
	Config         *config.Config
	Logger         logger.Interface
	Validation     *validator.Validate
	ContextTimeout time.Duration
	User           usecase.UserUsecase
	Doctor         usecase.DoctorUsecase
	Appointment    usecase.AppointmentUsecase
	Router         fiber.Router
}

func NewUserRoutes(c HandlerV1Config) {
	r := &HandlerV1{
		Config:         c.Config,
		Logger:         c.Logger,
		Validation:     c.Validation,
		ContextTimeout: c.ContextTimeout,
		User:           c.User,
		Doctor:         c.Doctor,
		Appointment:    c.Appointment,
		Router:         c.Router,
	}

	// Should not be role required in signup
	userGroup := r.Router.Group("/auth")
	{
		userGroup.Post("/signup", r.SignUpUser)
		userGroup.Post("/signin", r.SignInUser)
	}

	// Ping
	r.Router.Get("/ping", r.Ping)

}
