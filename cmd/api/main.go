package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mfarooqzahid/go-social/internal/config"
	"github.com/mfarooqzahid/go-social/internal/db"
	"github.com/mfarooqzahid/go-social/internal/router"
)

func main() {
	config.LoadEnv()

	conn, err := db.ConnectDb()
	if err != nil {
		log.Fatalf("🛑 Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
		},
	)

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	// registering routes
	router.RegisterRoutes(app)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	addr := config.Envs.PublicHost + ":" + config.Envs.Port

	go func() {
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	log.Println("🚀 Server is running: ", addr)

	<-stop
	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Could not gracefully shut down Fiber: %v", err)
	}

	log.Println("✅ Server shut down successfully")

}
