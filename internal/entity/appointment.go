package entity

import "time"

const StatusBooked = "scheduled"

type Appointment struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	DoctorID        int       `json:"doctor_id"`
	AppointmentTime time.Time `json:"appointment_time"`
	Duration        int       `json:"duration"` // in minutes
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
