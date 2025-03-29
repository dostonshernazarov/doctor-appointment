package persistent

import (
	"context"
	"fmt"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/pkg/postgres"
)

// AppointmentRepo -.
type AppointmentRepo struct {
	*postgres.Postgres
}

// NewUser -.
func NewAppointment(pg *postgres.Postgres) *AppointmentRepo {
	return &AppointmentRepo{pg}
}

// CreateAppointment -
func (r *AppointmentRepo) CreateAppointment(ctx context.Context, appointment entity.Appointment) error {
	sql, args, err := r.Builder.
		Select("id", "user_id", "doctor_id", "appointment_time", "duration", "status", "created_at", "updated_at").
		From("appointments").
		Where("doctor_id = ?", appointment.DoctorID).
		Where("appointment_time = ?", appointment.AppointmentTime).
		ToSql()

	if err != nil {
		return fmt.Errorf("AppointmentRepo - CreateAppointment - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var existingAppointment entity.Appointment
	err = row.Scan(&existingAppointment.ID, &existingAppointment.UserID, &existingAppointment.DoctorID, &existingAppointment.AppointmentTime, &existingAppointment.Duration, &existingAppointment.Status, &existingAppointment.CreatedAt, &existingAppointment.UpdatedAt)
	if err != nil {
		return fmt.Errorf("AppointmentRepo - CreateAppointment - row.Scan: %w", err)
	}

	if existingAppointment.ID != 0 {
		return fmt.Errorf("AppointmentRepo - CreateAppointment - appointment already booked")
	}

	sql, args, err = r.Builder.
		Insert("appointments").
		Columns("user_id", "doctor_id", "appointment_time", "duration", "status").
		Values(appointment.UserID, appointment.DoctorID, appointment.AppointmentTime, appointment.Duration, appointment.Status).ToSql()

	if err != nil {
		return fmt.Errorf("AppointmentRepo - CreateAppointment - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AppointmentRepo - CreateAppointment - r.Pool.Exec: %w", err)
	}

	return nil

}

// GetAppointmentByID -.
func (r *AppointmentRepo) GetAppointmentByID(ctx context.Context, id int) (entity.Appointment, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "doctor_id", "appointment_time", "duration", "status", "created_at", "updated_at").
		From("appointments").
		Where("id = ?", id).
		Limit(1).
		ToSql()

	if err != nil {
		return entity.Appointment{}, fmt.Errorf("AppointmentRepo - GetAppointmenrByID - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var appointment entity.Appointment
	err = row.Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.AppointmentTime, &appointment.Duration, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)
	if err != nil {
		return entity.Appointment{}, fmt.Errorf("AppointmentRepo - GetAppointmentByID - row.Scan: %w", err)
	}

	return appointment, nil

}

// GetAppointmentsByUserID -.
func (r *AppointmentRepo) GetAppointmentsByUserID(ctx context.Context, userID int) ([]entity.Appointment, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "doctor_id", "appointment_time", "duration", "status", "created_at", "updated_at").
		From("appointments").
		Where("user_id = ?", userID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetAppointmentsByUserID - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetAppointmentsByUserID - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var appointments []entity.Appointment
	for rows.Next() {
		var appointment entity.Appointment
		err = rows.Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.AppointmentTime, &appointment.Duration, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("AppointmentRepo - GetAppointmentsByUserID - rows.Scan: %w", err)
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// GetAppointmentsByDoctorID -.
func (r *AppointmentRepo) GetAppointmentsByDoctorID(ctx context.Context, doctorID int) ([]entity.Appointment, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "doctor_id", "appointment_time", "duration", "status", "created_at", "updated_at").
		From("appointments").
		Where("doctor_id = ?", doctorID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetAppointmentsByDoctorID - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetAppointmentsByDoctorID - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var appointments []entity.Appointment
	for rows.Next() {
		var appointment entity.Appointment
		err = rows.Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.AppointmentTime, &appointment.Duration, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("AppointmentRepo - GetAppointmentsByDoctorID - rows.Scan: %w", err)
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// UpdateAppointment -.
func (r *AppointmentRepo) UpdateAppointment(ctx context.Context, appointment entity.Appointment) error {
	sql, args, err := r.Builder.
		Update("appointments").
		Set("status", appointment.Status).
		Where("id = ?", appointment.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("AppointmentRepo - UpdateAppointment - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AppointmentRepo - UpdateAppointment - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteAppointment -.
func (r *AppointmentRepo) DeleteAppointment(ctx context.Context, id int) error {
	sql, args, err := r.Builder.Delete("appointments").Where("id = ?", id).ToSql()
	if err != nil {
		return fmt.Errorf("AppointmentRepo - DeleteAppointment - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AppointmentRepo - DeleteAppointment - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetBookedAppointmentsByDoctorId -.
func (r *AppointmentRepo) GetBookedAppointmentsByDoctorId(ctx context.Context, doctorID int) ([]entity.Appointment, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "doctor_id", "appointment_time", "duration", "status", "created_at", "updated_at").
		From("appointments").
		Where("doctor_id = ?", doctorID).
		Where("status = ?", entity.StatusBooked).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetBookedAppointmentsByDoctorId - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetBookedAppointmentsByDoctorId - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var appointments []entity.Appointment
	for rows.Next() {
		var appointment entity.Appointment
		err = rows.Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.AppointmentTime, &appointment.Duration, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("AppointmentRepo - GetBookedAppointmentsByDoctorId - rows.Scan: %w", err)
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// GetBookedAppointmentsByUserId -.
func (r *AppointmentRepo) GetBookedAppointmentsByUserId(ctx context.Context, userID int) ([]entity.Appointment, error) {
	sql, args, err := r.Builder.
		Select("id", "user_id", "doctor_id", "appointment_time", "duration", "status", "created_at", "updated_at").
		From("appointments").
		Where("user_id = ?", userID).
		Where("status = ?", entity.StatusBooked).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetBookedAppointmentsByUserId - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AppointmentRepo - GetBookedAppointmentsByUserId - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var appointments []entity.Appointment
	for rows.Next() {
		var appointment entity.Appointment
		err = rows.Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.AppointmentTime, &appointment.Duration, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("AppointmentRepo - GetBookedAppointmentsByUserId - rows.Scan: %w", err)
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
