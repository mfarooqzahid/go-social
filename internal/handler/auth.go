package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfarooqzahid/go-social/internal/models"
	"github.com/mfarooqzahid/go-social/internal/services"
	"github.com/mfarooqzahid/go-social/internal/utils"
)

func Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "invalid request body",
			},
		)
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "missing fields required",
			},
		)
	}

	if !utils.ValidateEmail(loginRequest.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "email is not valid",
			},
		)
	}

	res, err := services.LoginUser(c.Context(), loginRequest)

	if err.Error != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func Signup(c *fiber.Ctx) error {
	var signupRequest models.SignupRequest

	if err := c.BodyParser(&signupRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "invalid request body",
			},
		)
	}

	if signupRequest.Email == "" || signupRequest.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "missing fields required",
			},
		)
	}

	if !utils.ValidateEmail(signupRequest.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "email is not valid",
			},
		)
	}

	res, err := services.RegisterUser(c.Context(), signupRequest)

	if err.Error != nil {
		return c.Status(err.StatusCode).JSON(fiber.Map{
			"error": err.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
