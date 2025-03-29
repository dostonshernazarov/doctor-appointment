package models

import (
	"time"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
)

type Appointment struct {
	DoctorID        int           `json:"doctor_id"`
	UserID          int           `json:"user_id"`
	AppointmentTime time.Time     `json:"appointment_time"`
	Duration        time.Duration `json:"duration"`
	Status          string        `json:"status"`
}

type AppointmentResponse struct {
	ID       int `json:"id"`
	DoctorID int `json:"doctor_id"`
	UserID   int `json:"user_id"`
}

type AppointmentsResponse struct {
	Appointments []entity.Appointment `json:"appointments"`
}
