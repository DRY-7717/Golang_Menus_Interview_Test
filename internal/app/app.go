package app

import (
	"golang_menu_interview/config"
	"golang_menu_interview/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()

	app := router.Init(cfg)


	go func() {

		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}

		err := app.Listen(":" + cfg.App.AppPort)

		if err != nil {
			log.Error().Err(err).Msg("Error starting server")
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	<-quit

	log.Info().Msg("server shutdown of 5 second.")

	done := make(chan struct{})
	go func() {
		if err := app.Shutdown(); err != nil {
			log.Error().Err(err).Msg("error shutting down server")
		}
		close(done)
	}()

	select {
	case <-done:
		log.Info().Msg("server stopped gracefully")
	case <-time.After(5 * time.Second):
		log.Warn().Msg("forced shutdown after 5 seconds")
	}

}
