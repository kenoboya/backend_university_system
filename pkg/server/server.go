package server

import "time"

type ServerConfig struct {
	Addr           string        `yaml:"port"`
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"`
	ReadTimeout    time.Duration `yaml:"readTimeout"`
	WriteTimeout   time.Duration `yaml:"writeTimeout"`
}
