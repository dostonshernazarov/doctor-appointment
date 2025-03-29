package v1

import (
	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Ping
// @Description Ping the server
// @Accept json
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func (r *HandlerV1) Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(models.PingResponse{
		Message: "pong",
	})
}
