package config

import (
	"os"
	database "test-crud/pkg/database/psql"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	HTTP HTTPConfig
	PSQl database.PSQlConfig
	Auth AuthConfig
}
type HTTPConfig struct {
	Addr           string        `mapstructure:"port"`
	MaxHeaderBytes int           `mapstructure:"maxHeaderBytes"`
	ReadTimeout    time.Duration `mapstructure:"readTimeout"`
	WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
}
type AuthConfig struct {
	JWT          JWTConfig
	PasswordSalt string
}
type JWTConfig struct {
	AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
	RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	SecretKey       string
}

func Init(configDir string) (*Config, error) {
	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := setFromEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func setFromEnv(config *Config) error {
	if err := gotenv.Load("../../.env"); err != nil {
		return err
	}
	if err := envconfig.Process("DB", &config.PSQl); err != nil {
		return err
	}
	config.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	config.Auth.JWT.SecretKey = os.Getenv("SECRET_KEY")
	return nil
}
func unmarshal(config *Config) error {
	if err := viper.UnmarshalKey("http", &config.HTTP); err != nil {
		return err
	}
	// AUTH во время дебага
	if err := viper.UnmarshalKey("stage", &config.Auth.JWT); err != nil {
		return err
	}
	return nil
}
func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("server")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.MergeInConfig()
}
