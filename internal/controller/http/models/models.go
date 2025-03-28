package models

type SignUpUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type SignInUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone"`
}
