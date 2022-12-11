package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	memeConfig "github.com/zhews/memed-simple/pkg/config/meme"
	"github.com/zhews/memed-simple/pkg/handler"
	"github.com/zhews/memed-simple/pkg/handler/middleware"
	"github.com/zhews/memed-simple/pkg/repository/postgres"
	"github.com/zhews/memed-simple/pkg/service"
	"log"
)

func main() {
	config, err := memeConfig.ParseFromEnvironmentalVariables()
	if err != nil {
		log.Fatalf("Could not parse config: %s\n", err)
	}
	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
	}
	postgresRepository := &postgres.MemeRepositoryPostgres{
		Conn: pool,
	}
	memeService := service.MemeService{
		Config:     config,
		Repository: postgresRepository,
	}
	memeHandler := handler.MemeHandler{
		Config:  config,
		Service: memeService,
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.CorsAllowOrigins,
	}))
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS512",
		SigningKey:    []byte(config.AccessSecretKey),
	}))
	app.Get("/", memeHandler.HandleGetMemes)
	app.Get("/:id", memeHandler.HandleGetMeme)
	app.Post("/", memeHandler.HandleUploadMeme)
	app.Put("/:id", memeHandler.HandleUpdateMeme)
	app.Delete("/:id", middleware.AdminOnly, memeHandler.HandleDeleteMeme)
	app.Static("/image", config.MemeDirectory)
	app.Get("/health", handler.HandleHealth)
	log.Fatalf("Error while running the HTTP server: %s\n", app.Listen(fmt.Sprintf(":%d", config.Port)))
}
