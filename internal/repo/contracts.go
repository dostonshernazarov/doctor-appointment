package repo

import (
	"context"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
)

type (
	// UserRepo -.
	UserRepo interface {
		CreateUser(ctx context.Context, user entity.User) (int, error)
		GetUserByID(ctx context.Context, id int) (entity.User, error)
		GetUserByEmail(ctx context.Context, email string) (entity.User, error)
		ListUsers(ctx context.Context) ([]entity.User, error)
		UpdateUser(ctx context.Context, user entity.UserUpdate) error
		DeleteUser(ctx context.Context, id int) error
		GetPasswordHash(ctx context.Context, email string) (entity.GetPasswordHash, error)
		UpdateToken(ctx context.Context, id int, token string) error
	}

	// AppointmentRepo -.
	AppointmentRepo interface {
		CreateAppointment(ctx context.Context, appointment entity.Appointment) error
		GetAppointmentsByDoctorID(ctx context.Context, doctorID int) ([]entity.Appointment, error)
		GetBookedAppointmentsByDoctorId(ctx context.Context, doctorID int) ([]entity.Appointment, error)
		UpdateAppointment(ctx context.Context, appointment entity.Appointment) error
		DeleteAppointment(ctx context.Context, id int) error
		GetBookedAppointmentsByUserId(ctx context.Context, userID int) ([]entity.Appointment, error)
		GetAppointmentByID(ctx context.Context, id int) (entity.Appointment, error)
		GetAppointmentsByUserID(ctx context.Context, userID int) ([]entity.Appointment, error)
		GetAllAppointments(ctx context.Context) ([]entity.Appointment, error)
	}

	// DoctorRepo -.
	DoctorRepo interface {
		CreateDoctor(ctx context.Context, doctor entity.Doctor) error
		GetDoctorByID(ctx context.Context, id int) (entity.Doctor, error)
		GetDoctorBySpecialization(ctx context.Context, specialization string) ([]entity.Doctor, error)
		GetDoctors(ctx context.Context) ([]entity.Doctor, error)
		UpdateDoctor(ctx context.Context, doctor entity.Doctor) error
		DeleteDoctor(ctx context.Context, id int) error
		ListSpecializations(ctx context.Context) ([]string, error)
		GetBookedSchedulesByDoctorID(ctx context.Context, doctorID int) ([]entity.Schedule, error)
	}
)
