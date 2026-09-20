package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tronopay/config"
	"tronopay/internal/httpd"
	"tronopay/internal/repo"
	"tronopay/internal/service"
	"tronopay/pkg/postgres"

	"github.com/sirupsen/logrus"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("app - Run - config error: %v", err)
	}

	// Miniapp for health checking
	healthcheck(cfg.HTTP.Port)

	// Logger
	log := initLogger(cfg.Log.Level)
	log.Info("Configuration loaded successfully.")

	// DB
	log.Info("Init postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatalf("app - Run - postgres error: %v", err)
	}
	defer pg.Close()

	// Repositories
	log.Info("Init repositories...")
	repositories := repo.NewRepositories(pg)

	// Services
	log.Info("Init services...")
	services := service.NewServices(&service.ServicesDependencies{
		App:   &cfg.App,
		Log:   log,
		Repos: repositories,
	})

	// Web server UI
	log.Info("Init HTTP server...")
	srv := httpd.RunServer(cfg.HTTP.Port, services)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("STOP signal received")

	// Wait 5 sec
	log.Info("Waitting...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Server stopping
	log.Info("Server stopping...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Info("Server stopped")
}

func healthcheck(port string) {
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		resp, err := http.Get("http://localhost:" + port + "/health")
		if err == nil && resp.StatusCode == http.StatusOK {
			os.Exit(0)
		}
		os.Exit(1)
	}
}

func initLogger(logLevel string) *logrus.Logger {
	log := logrus.New()

	logrusLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrusLevel)
	}
	log.SetOutput(os.Stdout)

	return log
}
