package router

import (
	"golang_menu_interview/config"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
)

func Init(config *config.Config) *fiber.App {

	app := fiber.New(fiber.Config{
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
		},
	))

	var docs string

	if config.App.AppEnv == "production" {
		docs = "./docs/prod.swagger.json"
	} else {
		docs = "./docs/dev.swagger.json"
	}

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: docs,
		Path:     "swagger",
		Title:    "Swagger API Docs",
		CacheAge: -1,
	}))

	validator := validator.New()

	db, err := config.ConnectionPostgres()
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to database")
	}

	api := app.Group("/api")

	// check api run
	api.Get("/check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"message":   "Server is running",
			"timestamp": time.Now().Unix(),
		})
	})

	MenuRouter(api, db.DB, validator)

	return app

}
