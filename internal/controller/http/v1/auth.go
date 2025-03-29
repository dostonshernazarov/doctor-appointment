package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/models"
	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/pkg/etc"
	tokens "github.com/dostonshernazarov/doctor-appointment/pkg/token"
	"github.com/gofiber/fiber/v2"
)

// @Summary Sign up user
// @Description Sign up a user
// @Accept json
// @Produce json
// @Tags auth
// @Param user body models.SignUpUserRequest true "User"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /auth/signup [post]
func (r *HandlerV1) SignUpUser(c *fiber.Ctx) error {
	var req models.SignUpUserRequest
	if err := c.BodyParser(&req); err != nil {
		r.Logger.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := r.Validation.Struct(req); err != nil {
		r.Logger.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	hashedPassword, err := etc.HashPassword(req.Password)
	if err != nil {
		r.Logger.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusInternalServerError, "failed to hash password")
	}

	token, err := tokens.GenerateJWTToken(req.Email, string(entity.RoleUser), r.Config.Jwt.Secret, time.Duration(r.Config.Jwt.ExpiresAt)*time.Second)
	if err != nil {
		r.Logger.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusInternalServerError, "failed to generate token")
	}

	err = r.User.CreateUser(c.Context(), entity.User{
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
		Role:     entity.RoleUser,
		Token:    token,
	})

	if err != nil {
		r.Logger.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusInternalServerError, "failed to create user")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created successfully", "token": token})

}

// @Summary Sign in user
// @Description Sign in a user
// @Accept json
// @Produce json
// @Tags auth
// @Param user body models.SignInUserRequest true "User"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /auth/signin [post]
func (r *HandlerV1) SignInUser(c *fiber.Ctx) error {
	var req models.SignInUserRequest
	if err := c.BodyParser(&req); err != nil {
		r.Logger.Error(err, "http - v1 - sign in user")

		return errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	user, err := r.User.GetPasswordHash(c.Context(), req.Email)
	if err != nil {
		r.Logger.Error(err, "http - v1 - sign in user")

		return errorResponse(c, http.StatusInternalServerError, "failed to get password hash")
	}

	fmt.Println("password hash", user.PasswordHash, "Request password", req.Password)

	if !etc.CheckPasswordHash(req.Password, user.PasswordHash) {
		r.Logger.Error("invalid credentials", "http - v1 - sign in user")

		return errorResponse(c, http.StatusUnauthorized, "invalid credentials")
	}

	token, err := tokens.GenerateJWTToken(req.Email, string(entity.RoleUser), r.Config.Jwt.Secret, time.Duration(r.Config.Jwt.ExpiresAt)*time.Second)
	if err != nil {
		r.Logger.Error(err, "http - v1 - sign in user")

		return errorResponse(c, http.StatusInternalServerError, "failed to generate token")
	}

	err = r.User.UpdateToken(c.Context(), user.ID, token)
	if err != nil {
		r.Logger.Error(err, "http - v1 - sign in user")

		return errorResponse(c, http.StatusInternalServerError, "failed to update token")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User signed in successfully", "token": token})
}
