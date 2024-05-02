package app

import (
	"net/http"
	config "test-crud/internal/config"
	"test-crud/internal/repository/psql"
	"test-crud/internal/service"
	"test-crud/internal/transports/rest"
	"test-crud/pkg/auth"
	database "test-crud/pkg/database/psql"
	"test-crud/pkg/hash"
	"test-crud/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Run(configPath string) {
	log := logger.CreateLogger()
	zap.ReplaceGlobals(log)
	defer log.Sync()

	config, err := config.Init(configPath)
	if err != nil {
		logger.Fatal(
			zap.String("package", "internal/app"),
			zap.String("file", "app.go"),
			zap.String("function", "config.Init()"),
			zap.Error(err),
		)
	}

	db := connectToDatabase(*config)
	defer db.Close()

	hasher := hash.NewSHA256Hasher(config.Auth.PasswordSalt)
	tokenManager, err := auth.NewManager(config.Auth.JWT.SecretKey)
	if err != nil {
		logger.Fatal(
			zap.String("package", "internal/app"),
			zap.String("file", "app.go"),
			zap.String("function", "srv.ListenAndServe()"),
			zap.Error(err),
		)
	}

	repositories := psql.NewRepositories(db)
	deps := service.Deps{
		Repos:           *repositories,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		AccessTokenTTL:  config.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: config.Auth.JWT.RefreshTokenTTL,
	}
	services := service.NewServices(deps)
	handler := rest.NewHandler(services, *tokenManager)

	srv := &http.Server{
		Addr:           config.HTTP.Addr,
		ReadTimeout:    config.HTTP.ReadTimeout,
		WriteTimeout:   config.HTTP.WriteTimeout,
		MaxHeaderBytes: config.HTTP.MaxHeaderBytes,
		Handler:        handler.InitRouter(),
	}
	logger.Infof("Server started at ")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(
			zap.String("package", "internal/app"),
			zap.String("file", "app.go"),
			zap.String("function", "srv.ListenAndServe()"),
			zap.Error(err),
		)
	}
}
func connectToDatabase(cfg config.Config) *sqlx.DB {
	db, err := database.NewPostgresConnection(cfg.PSQl)
	if err != nil {
		logger.Fatal(
			zap.String("package", "internal/app"),
			zap.String("file", "app.go"),
			zap.String("function", "connectToDatabase()"),
			zap.Error(err),
		)
	}
	return db
}
