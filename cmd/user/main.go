package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	userConfig "github.com/zhews/memed-simple/pkg/config/user"
	"github.com/zhews/memed-simple/pkg/handler"
	"github.com/zhews/memed-simple/pkg/repository/postgres"
	"github.com/zhews/memed-simple/pkg/service"
	"log"
)

func main() {
	config, err := userConfig.FromEnvironmentalVariables()
	if err != nil {
		log.Fatalf("Could not parse config: %s\n", err)
	}
	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
	}
	postgresRepository := postgres.UserRepositoryPostgres{
		Conn: pool,
	}
	userService := service.UserService{
		Config:     config,
		Repository: &postgresRepository,
	}
	userHandler := handler.UserHandler{
		Config:  config,
		Service: userService,
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.CorsAllowOrigins,
	}))
	app.Get("/health", handler.HandleHealth)
	app.Get("/:id", userHandler.HandleGetById)
	auth := app.Group("/auth")
	auth.Post("/register", userHandler.HandleRegister)
	auth.Post("/login", userHandler.HandleLogin)
	auth.Get("/check/username/:username", userHandler.HandleCheckUsername)
	auth.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS512",
		SigningKey:    []byte(config.RefreshSecretKey),
		TokenLookup:   "cookie:refreshToken",
	}))
	auth.Post("/refresh", userHandler.HandleRefresh)
	auth.Post("/logout", userHandler.HandleLogout)
	log.Fatalf("Error while running the HTTP server: %s\n", app.Listen(fmt.Sprintf(":%d", config.Port)))
}
