package persistent

import (
	"context"
	"fmt"
	"time"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/pkg/postgres"
)

// DoctorRepo -.
type DoctorRepo struct {
	*postgres.Postgres
}

// NewUser -.
func NewDoctor(pg *postgres.Postgres) *DoctorRepo {
	return &DoctorRepo{pg}
}

// CreateDoctor - creates a new doctor in the database.
func (r *DoctorRepo) CreateDoctor(ctx context.Context, doctor entity.Doctor) error {
	sql, args, err := r.Builder.
		Insert("doctors").
		Columns("name", "specialization", "schedule").
		Values(doctor.Name, doctor.Specialization, doctor.Schedule).ToSql()

	if err != nil {
		return fmt.Errorf("DoctorRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("DoctorRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetDoctorByID -.
func (r *DoctorRepo) GetDoctorByID(ctx context.Context, id int) (entity.Doctor, error) {
	sql, args, err := r.Builder.
		Select("id", "name", "specialization", "schedule", "created_at", "updated_at").
		From("doctors").
		Where("id = ?", id).
		Limit(1).
		ToSql()

	if err != nil {
		return entity.Doctor{}, fmt.Errorf("DoctorRepo - GetDoctorByID - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var doctor entity.Doctor
	err = row.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.Schedule, &doctor.CreatedAt, &doctor.UpdatedAt)
	if err != nil {
		return entity.Doctor{}, fmt.Errorf("DoctorRepo - GetDoctorByID - row.Scan: %w", err)
	}

	return doctor, nil
}

// GetDoctors -.
func (r *DoctorRepo) GetDoctors(ctx context.Context) ([]entity.Doctor, error) {
	sql, args, err := r.Builder.
		Select("id", "name", "specialization", "schedule", "created_at", "updated_at").
		From("doctors").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - GetDoctors - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - GetDoctors - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var doctors []entity.Doctor
	for rows.Next() {
		var doctor entity.Doctor
		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.Schedule, &doctor.CreatedAt, &doctor.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("DoctorRepo - GetDoctors - rows.Scan: %w", err)
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

// UpdateDoctor -.
func (r *DoctorRepo) UpdateDoctor(ctx context.Context, doctor entity.Doctor) error {
	updateTime := time.Now()

	sql, args, err := r.Builder.
		Update("doctors").
		Set("name", doctor.Name).
		Set("specialization", doctor.Specialization).
		Set("schedule", doctor.Schedule).
		Set("updated_at", updateTime).
		Where("id = ?", doctor.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("DoctorRepo - UpdateDoctor - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("DoctorRepo - UpdateDoctor - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteDoctor -.
func (r *DoctorRepo) DeleteDoctor(ctx context.Context, id int) error {
	sql, args, err := r.Builder.
		Delete("doctors").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return fmt.Errorf("DoctorRepo - DeleteDoctor - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("DoctorRepo - DeleteDoctor - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetDoctorBySpecialization -.
func (r *DoctorRepo) GetDoctorBySpecialization(ctx context.Context, specialization string) ([]entity.Doctor, error) {
	sql, args, err := r.Builder.
		Select("id", "name", "specialization", "schedule", "created_at", "updated_at").
		From("doctors").
		Where("specialization = ?", specialization).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - GetDoctorBySpecialization - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - GetDoctorBySpecialization - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var doctors []entity.Doctor
	for rows.Next() {
		var doctor entity.Doctor
		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.Schedule, &doctor.CreatedAt, &doctor.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("DoctorRepo - GetDoctorBySpecialization - rows.Scan: %w", err)
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

// ListSpecializations -.
func (r *DoctorRepo) ListSpecializations(ctx context.Context) ([]string, error) {
	sql, args, err := r.Builder.
		Select("specialization").
		From("doctors").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - ListSpecializations - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - ListSpecializations - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var specializations []string
	for rows.Next() {
		var specialization string
		err = rows.Scan(&specialization)
		if err != nil {
			return nil, fmt.Errorf("DoctorRepo - ListSpecializations - rows.Scan: %w", err)
		}

		specializations = append(specializations, specialization)
	}

	return specializations, nil
}

// GetBookedSchedulesByDoctorID -.
func (r *DoctorRepo) GetBookedSchedulesByDoctorID(ctx context.Context, doctorID int) ([]entity.Schedule, error) {
	sql, args, err := r.Builder.
		Select("schedule").
		From("doctors").
		Where("id = ?", doctorID).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - GetBookedSchedulesByDoctorID - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("DoctorRepo - GetBookedSchedulesByDoctorID - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var schedules []entity.Schedule
	for rows.Next() {
		var schedule entity.Schedule
		err = rows.Scan(&schedule)
		if err != nil {
			return nil, fmt.Errorf("DoctorRepo - GetBookedSchedulesByDoctorID - rows.Scan: %w", err)
		}

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}
