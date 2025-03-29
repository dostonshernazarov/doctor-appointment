package models

import (
	"time"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
)

type Doctor struct {
	Name           string   `json:"name"`
	Specialization string   `json:"specialization"`
	Schedule       Schedule `json:"schedule"`
}

type Schedule struct {
	Days  []string `json:"days"`
	Start string   `json:"start"`
	End   string   `json:"end"`
}

type DoctorResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Specialization string    `json:"specialization"`
	Schedule       Schedule  `json:"schedule"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type AllDoctorsResponse struct {
	Doctors []entity.Doctor `json:"doctors"`
}

type SpecializationResponse struct {
	Specializations []string `json:"specializations"`
}

type BookedSchedulesResponse struct {
	BookedSchedules []entity.Schedule `json:"booked_schedules"`
}
