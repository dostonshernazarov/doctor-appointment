package usecase

import (
	"context"
)

type (
	// UserUsecase -.
	UserUsecase interface {
		CreateUser(ctx context.Context, user User) error
		GetUserByID(ctx context.Context, id int) (User, error)
		GetUserByEmail(ctx context.Context, email string) (User, error)
		ListUsers(ctx context.Context) ([]User, error)
		UpdateUser(ctx context.Context, user User) error
		DeleteUser(ctx context.Context, id int) error
	}

	// AppointmentUsecase -.
	AppointmentUsecase interface {
		CreateAppointment(ctx context.Context, appointment Appointment) error
		GetAppointmentsByDoctorID(ctx context.Context, doctorID int) ([]Appointment, error)
		GetBookedAppointmentsByDoctorId(ctx context.Context, doctorID int) ([]Appointment, error)
		UpdateAppointment(ctx context.Context, appointment Appointment) error
		DeleteAppointment(ctx context.Context, id int) error
		GetBookedAppointmentsByUserId(ctx context.Context, userID int) ([]Appointment, error)
		GetAppointmentByID(ctx context.Context, id int) (Appointment, error)
		GetAppointmentsByUserID(ctx context.Context, userID int) ([]Appointment, error)
	}

	// DoctorUsecase -.
	DoctorUsecase interface {
		CreateDoctor(ctx context.Context, doctor Doctor) error
		GetDoctorByID(ctx context.Context, id int) (Doctor, error)
		GetDoctorByEmail(ctx context.Context, email string) (Doctor, error)
		ListDoctors(ctx context.Context) ([]Doctor, error)
		UpdateDoctor(ctx context.Context, doctor Doctor) error
		DeleteDoctor(ctx context.Context, id int) error
		ListSpecializations(ctx context.Context) ([]string, error)
	}
)
