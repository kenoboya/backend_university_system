package app

import (
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
	db := connectToDatabase(configPath)
	defer db.Close()
	repositories := psql.NewRepositories(*psql.NewStudents(db))
	services := service.NewServices(*service.NewStudents(&repositories.Students))
	handler := rest.NewHandler(services)

	serverConfig, err := config.ReadServerConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:           serverConfig.Addr,
		ReadTimeout:    serverConfig.ReadTimeout,
		WriteTimeout:   serverConfig.WriteTimeout,
		MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		Handler:        handler.InitRouter(),
	}
	log.Println("Server started at ", time.Now().Format(time.RFC3339))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
func connectToDatabase(configPath string) *sqlx.DB {
	configDataBase, err := config.ReadDatabaseConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.NewPostgresConnection(configDataBase)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
