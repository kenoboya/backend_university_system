package config

import (
	"os"
	"path/filepath"
	database "test-crud/pkg/database/psql"
	"test-crud/pkg/server"

	"gopkg.in/yaml.v2"
)

func ReadServerConfig(configDir string) (server.ServerConfig, error) {
	filepath := filepath.Join(configDir, "server.yml")
	var config server.ServerConfig
	content, err := os.ReadFile(filepath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
func ReadDatabaseConfig(configDir string) (database.PSQlConfig, error) {
	filepath := filepath.Join(configDir, "database.yml")
	var config database.PSQlConfig
	content, err := os.ReadFile(filepath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
