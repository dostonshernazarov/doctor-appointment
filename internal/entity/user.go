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
	Token     string    `json:"token"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password_hash"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password_hash"`
	Phone    string `json:"phone"`
}
