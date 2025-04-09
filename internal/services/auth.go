package services

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/mfarooqzahid/go-social/internal/db"
	"github.com/mfarooqzahid/go-social/internal/models"
	"github.com/mfarooqzahid/go-social/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx context.Context, req models.LoginRequest) (models.AuthResponse, error) {
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
	if err != nil || err == sql.ErrNoRows {

		return authResponse, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return authResponse, errors.New("Invalid email or password: " + err.Error())
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return authResponse, errors.New("error generating token")
	}

	authResponse = models.AuthResponse{
		Token: token,
		User:  user,
	}

	return authResponse, nil
}
