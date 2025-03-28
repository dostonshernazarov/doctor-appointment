package common

import (
	"context"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/internal/repo"
)

type UseCase struct {
	userRepo        repo.UserRepo
	doctorRepo      repo.DoctorRepo
	appointmentRepo repo.AppointmentRepo
}

func NewUseCase(userRepo repo.UserRepo, doctorRepo repo.DoctorRepo, appointmentRepo repo.AppointmentRepo) *UseCase {
	return &UseCase{
		userRepo:        userRepo,
		doctorRepo:      doctorRepo,
		appointmentRepo: appointmentRepo,
	}
}

// CreateUser -.
func (uc *UseCase) CreateUser(ctx context.Context, user entity.User) error {
	return uc.userRepo.CreateUser(ctx, user)
}

// GetUserByID -.
func (uc *UseCase) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	return uc.userRepo.GetUserByID(ctx, id)
}

// GetUserByEmail -.
func (uc *UseCase) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return uc.userRepo.GetUserByEmail(ctx, email)
}

// ListUsers -.
func (uc *UseCase) ListUsers(ctx context.Context) ([]entity.User, error) {
	return uc.userRepo.ListUsers(ctx)
}

// UpdateUser -.
func (uc *UseCase) UpdateUser(ctx context.Context, user entity.User) error {
	return uc.userRepo.UpdateUser(ctx, user)
}

// DeleteUser -.
func (uc *UseCase) DeleteUser(ctx context.Context, id int) error {
	return uc.userRepo.DeleteUser(ctx, id)
}

// CreateDoctor -.
func (uc *UseCase) CreateDoctor(ctx context.Context, doctor entity.Doctor) error {
	return uc.doctorRepo.CreateDoctor(ctx, doctor)
}

// GetDoctorByID -.
func (uc *UseCase) GetDoctorByID(ctx context.Context, id int) (entity.Doctor, error) {
	return uc.doctorRepo.GetDoctorByID(ctx, id)
}

// GetDoctorByEmail -.
func (uc *UseCase) GetDoctorByEmail(ctx context.Context, email string) (entity.Doctor, error) {
	return uc.doctorRepo.GetDoctorByEmail(ctx, email)
}

// ListDoctors -.
func (uc *UseCase) ListDoctors(ctx context.Context) ([]entity.Doctor, error) {
	return uc.doctorRepo.ListDoctors(ctx)
}

// UpdateDoctor -.
func (uc *UseCase) UpdateDoctor(ctx context.Context, doctor entity.Doctor) error {
	return uc.doctorRepo.UpdateDoctor(ctx, doctor)
}

// DeleteDoctor -.
func (uc *UseCase) DeleteDoctor(ctx context.Context, id int) error {
	return uc.doctorRepo.DeleteDoctor(ctx, id)
}

// ListSpecializations -.
func (uc *UseCase) ListSpecializations(ctx context.Context) ([]string, error) {
	return uc.doctorRepo.ListSpecializations(ctx)
}

// CreateAppointment -.
func (uc *UseCase) CreateAppointment(ctx context.Context, appointment entity.Appointment) error {
	return uc.appointmentRepo.CreateAppointment(ctx, appointment)
}

// GetAppointmentsByDoctorID -.
func (uc *UseCase) GetAppointmentsByDoctorID(ctx context.Context, doctorID int) ([]entity.Appointment, error) {
	return uc.appointmentRepo.GetAppointmentsByDoctorID(ctx, doctorID)
}

// GetBookedAppointmentsByDoctorId -.
func (uc *UseCase) GetBookedAppointmentsByDoctorId(ctx context.Context, doctorID int) ([]entity.Appointment, error) {
	return uc.appointmentRepo.GetBookedAppointmentsByDoctorId(ctx, doctorID)
}

// UpdateAppointment -.
func (uc *UseCase) UpdateAppointment(ctx context.Context, appointment entity.Appointment) error {
	return uc.appointmentRepo.UpdateAppointment(ctx, appointment)
}

// DeleteAppointment -.
func (uc *UseCase) DeleteAppointment(ctx context.Context, id int) error {
	return uc.appointmentRepo.DeleteAppointment(ctx, id)
}

// GetBookedAppointmentsByUserId -.
func (uc *UseCase) GetBookedAppointmentsByUserId(ctx context.Context, userID int) ([]entity.Appointment, error) {
	return uc.appointmentRepo.GetBookedAppointmentsByUserId(ctx, userID)
}

// GetAppointmentByID -.
func (uc *UseCase) GetAppointmentByID(ctx context.Context, id int) (entity.Appointment, error) {
	return uc.appointmentRepo.GetAppointmentByID(ctx, id)
}

// GetAppointmentsByUserID -.
func (uc *UseCase) GetAppointmentsByUserID(ctx context.Context, userID int) ([]entity.Appointment, error) {
	return uc.appointmentRepo.GetAppointmentsByUserID(ctx, userID)
}
