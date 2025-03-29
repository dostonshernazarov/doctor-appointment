package v1

import (
	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/models"
	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/pkg/etc"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create user
// @Description Create user
// @Accept json
// @Produce json
// @Tags user
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users [post]
func (h *HandlerV1) CreateUser(c *fiber.Ctx) error {
	user := models.CreateUserRequest{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := etc.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := h.User.CreateUser(c.Context(), entity.User{
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
		Password: hashedPassword,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{
		ID:       id,
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
	})
}

// @Summary Get user by id
// @Description Get user by id
// @Accept json
// @Produce json
// @Tags user
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users/{id} [get]
func (h *HandlerV1) GetUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user, err := h.User.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
	})
}

// @Summary Update user
// @Description Update user
// @Accept json
// @Produce json
// @Tags user
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users/{id} [put]
func (h *HandlerV1) UpdateUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user := models.UpdateUserRequest{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.User.UpdateUser(c.Context(), entity.UserUpdate{
		ID:       userID,
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.UserResponse{
		ID:       userID,
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
	})

}

// @Summary Delete user
// @Description Delete user
// @Accept json
// @Produce json
// @Tags user
// @Param id path int true "User ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users/{id} [delete]
func (h *HandlerV1) DeleteUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	err := h.User.DeleteUser(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Message: "User deleted successfully",
	})
}

// @Summary Get all users
// @Description Get all users
// @Accept json
// @Produce json
// @Tags user
// @Success 200 {object} models.ListUsersResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users [get]
func (h *HandlerV1) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.User.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.ListUsersResponse{
		Users: users,
	})
}
