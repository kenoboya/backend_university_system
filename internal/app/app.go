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

	hasher := hash.NewSHA1Hasher(config.AuthConfig.PasswordSalt)
	tokenManager, err := auth.NewManager(config.AuthConfig.JWT.SecretKey)
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
		Repos:        *repositories,
		Hasher:       hasher,
		TokenManager: tokenManager,
	}
	services := service.NewServices(deps)
	handler := rest.NewHandler(services)

	srv := &http.Server{
		Addr:           config.ServerConfig.Addr,
		ReadTimeout:    config.ServerConfig.ReadTimeout,
		WriteTimeout:   config.ServerConfig.WriteTimeout,
		MaxHeaderBytes: config.ServerConfig.MaxHeaderBytes,
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
	db, err := database.NewPostgresConnection(cfg.PSQlConfig)
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
