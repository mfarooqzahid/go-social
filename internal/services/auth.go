package services

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mfarooqzahid/go-social/internal/db"
	"github.com/mfarooqzahid/go-social/internal/models"
	"github.com/mfarooqzahid/go-social/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx context.Context, req models.LoginRequest) (models.AuthResponse, models.Error) {
	var authResponse models.AuthResponse

	var user models.User

	row := db.PGX.QueryRow(
		ctx,
		`SELECT id, name, email, password, created_at, updated_at FROM users WHERE email=$1`,
		req.Email,
	)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		if err == sql.ErrNoRows {
			return authResponse, models.Error{
				Error:      errors.New("invalid email or password"),
				StatusCode: fiber.StatusUnauthorized, // 401 for authentication failure
			}
		}
		return authResponse, models.Error{
			Error:      errors.New("database error"),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return authResponse, models.Error{
			Error:      errors.New("invalid email or password"),
			StatusCode: fiber.StatusUnauthorized, // 401 for authentication failure
		}
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return authResponse, models.Error{
			Error:      errors.New("failed to generate authentication token"),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	authResponse = models.AuthResponse{
		Token: token,
		User:  user,
	}

	// Return empty models.Error to indicate success
	return authResponse, models.Error{}
}

func RegisterUser(ctx context.Context, req models.SignupRequest) (models.AuthResponse, models.Error) {
	var authResponse models.AuthResponse

	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := db.PGX.QueryRow(ctx, checkQuery, req.Email).Scan(&exists)
	if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return authResponse, models.Error{
			Error:      errors.New("failed to check user existence"),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	if exists {
		return authResponse, models.Error{
			Error:      errors.New("user with this email already exists"),
			StatusCode: fiber.StatusConflict, // 409 Conflict
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return authResponse, models.Error{
			Error:      errors.New("failed to process password"),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, created_at, updated_at
	`

	var user models.User

	err = db.PGX.QueryRow(
		ctx,
		query,
		req.Name,
		req.Email,
		string(hashedPassword),
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return authResponse, models.Error{
			Error:      errors.New("failed to create user"),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return authResponse, models.Error{
			Error:      errors.New("failed to generate authentication token"),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	authResponse = models.AuthResponse{
		Token: token,
		User:  user,
	}

	// Return empty models.Error to indicate success
	return authResponse, models.Error{}
}