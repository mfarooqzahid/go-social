package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/mfarooqzahid/go-social/internal/config"
)

var PGX *pgx.Conn

func ConnectDb() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.Envs.DBAddress)
	if err != nil {
		log.Fatalf("🛑 Unable to connect to database: %v", err)
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
		return nil, err
	}

	PGX = conn

	log.Println("✅ Successfully conected to database!")

	return conn, nil
}
