package v1

import (
	"net/http"

	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/models"
	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/internal/usecase"
	"github.com/dostonshernazarov/doctor-appointment/pkg/etc"
	"github.com/dostonshernazarov/doctor-appointment/pkg/logger"
	tokens "github.com/dostonshernazarov/doctor-appointment/pkg/token"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type userRoutes struct {
	u  usecase.UserUsecase
	l  logger.Interface
	v  *validator.Validate
	ts tokens.JwtHandler
}

func NewUserRoutes(apiV1Group fiber.Router, uc usecase.UserUsecase, l logger.Interface) {
	r := &userRoutes{
		u: uc,
		l: l,
		v: validator.New(),
	}

	userGroup := apiV1Group.Group("/users")
	{
		// userGroup.Post("/", r.CreateUser)
		// userGroup.Get("/:id", r.GetUserByID)
		// userGroup.Put("/:id", r.UpdateUser)
		// userGroup.Delete("/:id", r.DeleteUser)
	}

}

// @Summary Create user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.SignUpUserRequest true "User"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users/signup [post]
func (r *userRoutes) SignUpUser(c *fiber.Ctx) error {
	var req models.SignUpUserRequest
	if err := c.BodyParser(&req); err != nil {
		r.l.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	hashedPassword, err := etc.HashPassword(req.Password)
	if err != nil {
		r.l.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusInternalServerError, "failed to hash password")
	}

	err = r.u.CreateUser(c.Context(), entity.User{
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	})

	if err != nil {
		r.l.Error(err, "http - v1 - sign up user")

		return errorResponse(c, http.StatusInternalServerError, "failed to create user")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})

}

// @Summary Sign in user
// @Description Sign in a user
// @Accept json
// @Produce json
// @Param user body models.SignInUserRequest true "User"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users/signin [post]
func (r *userRoutes) SignInUser(c *fiber.Ctx) error {
	return nil
}

// @Summary Create user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users [post]
func (r *userRoutes) CreateUser(c *fiber.Ctx) error {
	return nil
}

// @Summary Get user by ID
// @Description Get a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users/{id} [get]
func (r *userRoutes) GetUserByID(c *fiber.Ctx) error {
	return nil
}

// @Summary Get user by email
// @Description Get a user by email
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users/email/{email} [get]
func (r *userRoutes) GetUserByEmail(c *fiber.Ctx) error {
	return nil
}

// @Summary List users
// @Description List all users
// @Accept json
// @Produce json
// @Success 200 {array} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users [get]
func (r *userRoutes) ListUsers(c *fiber.Ctx) error {
	return nil
}

// @Summary Update user
// @Description Update a user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users/{id} [put]
func (r *userRoutes) UpdateUser(c *fiber.Ctx) error {
	return nil
}

// @Summary Delete user
// @Description Delete a user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Router /users/{id} [delete]
func (r *userRoutes) DeleteUser(c *fiber.Ctx) error {
	return nil
}
