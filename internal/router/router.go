package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfarooqzahid/go-social/internal/handler"
)

func RegisterRoutes(app *fiber.App) {
	apiv1 := app.Group("/api")

	auth := apiv1.Group("/auth")
	{
		auth.Post("/login", handler.Login)
		auth.Post("/signup", handler.Signup)
	}

}
