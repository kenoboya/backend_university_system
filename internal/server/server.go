package server

import (
	"context"
	"net/http"
	"test-crud/internal/config"
	"test-crud/internal/transports/rest"
	"test-crud/pkg/logger"

	"go.uber.org/zap"
)

type Server struct {
	server *http.Server
}

func NewServer(config *config.Config, handler rest.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           config.HTTP.Addr,
			ReadTimeout:    config.HTTP.ReadTimeout,
			WriteTimeout:   config.HTTP.WriteTimeout,
			MaxHeaderBytes: config.HTTP.MaxHeaderBytes,
			Handler:        handler.InitRouter(),
		},
	}
}
func (s *Server) Run() error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	logger.Infof("The server started working ")
	return nil
}
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Fatal(
			zap.String("package", "internal/app"),
			zap.String("file", "server.go"),
			zap.String("function", "Shutdown()"),
			zap.Error(err),
		)
	}
	logger.Infof("The server finished working ")
	return nil
}
