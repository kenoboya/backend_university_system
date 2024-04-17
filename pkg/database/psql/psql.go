package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PSQlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	DBname   string `yaml:"dbname"`
	SSLmode  string `yaml:"sslmode"`
	Password string `yaml:"password"`
}

func NewPostgresConnection(inf PSQlConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		inf.Host, inf.Port, inf.Username, inf.DBname, inf.SSLmode, inf.Password))
	if err != nil {
		return nil, err
	}
	return db, nil
}
