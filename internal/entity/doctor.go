package entity

import "time"

type Doctor struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Specialization string    `json:"specialization"`
	Schedule       Schedule  `json:"schedule"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Schedule struct {
	Days  []string `json:"days"`
	Start string   `json:"start"`
	End   string   `json:"end"`
}
