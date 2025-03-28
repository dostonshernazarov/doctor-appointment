package entity

import (
	"time"
)

// User -.
type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password_hash"`
}
