package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/response"
	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	tokens "github.com/dostonshernazarov/doctor-appointment/pkg/token"
	"github.com/gofiber/fiber/v2"
)

const (
	UserKey   = "user"
	RoleKey   = "role"
	ClaimsKey = "claims"
)

type AuthConfig struct {
	Skipper   func(c *fiber.Ctx) bool
	JWTSecret string
}

func Authentication(config AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Skipper != nil && config.Skipper(c) {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "missing authorization header")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return response.ErrorResponse(c, http.StatusUnauthorized, "invalid authorization format")
		}

		claims, err := tokens.ParseToken(tokenParts[1], config.JWTSecret)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
		}

		// Add claims to context
		c.Locals(ClaimsKey, claims)

		// Create custom context
		ctx := context.WithValue(c.Context(), UserKey, claims.Email)
		ctx = context.WithValue(ctx, RoleKey, claims.Role)
		c.SetUserContext(ctx)

		return c.Next()
	}
}

// Role-based access control middleware
func RequireRole(requiredRole entity.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		userRole := ctx.Value(RoleKey).(entity.Role)

		if userRole != requiredRole {
			return response.ErrorResponse(c, http.StatusForbidden,
				"you don't have permission to access this resource")
		}

		return c.Next()
	}
}
