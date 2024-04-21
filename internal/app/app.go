package app

import (
	"fmt"
	"log"

	"net/http"
	config "test-crud/internal/config"
	"test-crud/internal/repository/psql"
	"test-crud/internal/service"
	"test-crud/internal/transports/rest"
	database "test-crud/pkg/database/psql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Run(configPath string) {
	config, err := config.Init(configPath)
	fmt.Println(config)
	if err != nil {
		log.Fatal(err)
	}
	db := connectToDatabase(*config)
	defer db.Close()
	repositories := psql.NewRepositories(*psql.NewStudents(db))
	services := service.NewServices(*service.NewStudents(&repositories.Students))
	handler := rest.NewHandler(services)

	srv := &http.Server{
		Addr:           config.ServerConfig.Addr,
		ReadTimeout:    config.ServerConfig.ReadTimeout,
		WriteTimeout:   config.ServerConfig.WriteTimeout,
		MaxHeaderBytes: config.ServerConfig.MaxHeaderBytes,
		Handler:        handler.InitRouter(),
	}
	log.Println("Server started at ", time.Now().Format(time.RFC3339))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
func connectToDatabase(cfg config.Config) *sqlx.DB {
	db, err := database.NewPostgresConnection(cfg.PSQlConfig)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
