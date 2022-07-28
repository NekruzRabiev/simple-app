package app

import (
	"context"
	"errors"
	"github.com/nekruzrabiev/simple-app/internal/config"
	delivery "github.com/nekruzrabiev/simple-app/internal/delivery/http"
	"github.com/nekruzrabiev/simple-app/internal/repository"
	"github.com/nekruzrabiev/simple-app/internal/server"
	"github.com/nekruzrabiev/simple-app/internal/service"
	"github.com/nekruzrabiev/simple-app/pkg/jwt"
	"github.com/nekruzrabiev/simple-app/pkg/logger"
	"github.com/nekruzrabiev/simple-app/pkg/rnd"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

func GetEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env != "" {
		return env
	}
	return "dev"
}

func Run(configPath string) {
	//Initialize configs
	cfg, err := config.Init(configPath, GetEnvironment(), ".env")
	if err != nil {
		logger.Errorf("failed to initialize config %v", err)
		return
	}

	dbConfig := &repository.ConfigPostgres{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	}
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		logger.Errorf("failed to init db: %v", err)
		return
	}

	repos := repository.NewRepositories(db)

	jwtManager, err := jwt.NewManager(cfg.Jwt.Salt, rnd.NewGeneratorRand(), 32)
	if err != nil {
		logger.Errorf("failed to init jwtManager: %v", err)
		return
	}

	reValidPassword := regexp.MustCompile(`^(([a-zA-ZА-Яа-я]+\d+)|(\d+[a-zA-ZА-Яа-я]+))[a-zA-zА-Яа-я\d]*$`)

	//TODO: change AccessTokenTTL and RefreshTokenTTL later
	deps := service.Deps{
		ReValidPassword: reValidPassword,
		Repos:           repos,
		JwtManager:      jwtManager,
		RndGen:          rnd.NewGeneratorRand(),
		AccessTokenTTL:  20 * time.Minute,
		RefreshTokenTTL: 365 * 24 * time.Hour,
	}

	services := service.NewServices(deps)

	handlers := delivery.NewHandler(services, jwtManager)

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := db.Close(); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
