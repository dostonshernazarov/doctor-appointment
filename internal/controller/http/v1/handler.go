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
	authGroup := r.Router.Group("/auth")
	{
		authGroup.Post("/signup", r.SignUpUser)
		authGroup.Post("/signin", r.SignInUser)
	}

	userGroup := r.Router.Group("/users")
	{
		userGroup.Post("/", r.CreateUser)
		userGroup.Get("/", r.GetAllUsers)
		userGroup.Get("/:id", r.GetUser)
		userGroup.Put("/:id", r.UpdateUser)
		userGroup.Delete("/:id", r.DeleteUser)
	}

	doctorGroup := r.Router.Group("/doctors")
	{
		doctorGroup.Post("/", r.CreateDoctor)
		doctorGroup.Get("/", r.GetAllDoctors)
		doctorGroup.Get("/:id", r.GetDoctorByID)
		doctorGroup.Put("/:id", r.UpdateDoctor)
		doctorGroup.Delete("/:id", r.DeleteDoctor)
		doctorGroup.Get("/specializations", r.ListSpecializations)
		doctorGroup.Get("/specialization/:specialization", r.GetDoctorsBySpecialization)
	}

	appointmentGroup := r.Router.Group("/appointments")
	{
		appointmentGroup.Post("/", r.CreateAppointment)
		appointmentGroup.Get("/:id", r.GetAppointmentByID)
		appointmentGroup.Put("/:id", r.UpdateAppointment)
		appointmentGroup.Delete("/:id", r.DeleteAppointment)
		appointmentGroup.Get("/doctor/:doctor_id", r.GetAppointmentsByDoctorID)
		appointmentGroup.Get("/user/:user_id", r.GetAppointmentsByUserID)
		appointmentGroup.Get("/doctor/:doctor_id/booked-schedules", r.GetBookedSchedulesByDoctorID)
		appointmentGroup.Get("/user/:user_id/booked-schedules", r.GetBookedSchedulesByUserID)
	}

	// Ping
	r.Router.Get("/ping", r.Ping)

}
