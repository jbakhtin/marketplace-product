package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	router "github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest/router/chi"
	"github.com/jbakhtin/marketplace-product/internal/modules/product"
	"github.com/jbakhtin/marketplace-product/internal/modules/product/ports"
)

type Server struct {
	logger ports.Logger
	http.Server
}

type Config interface {
	GetWebServerRestAddress() string
	GetAppKey() string
}

func NewWebServer(
	cfg Config,
	logger ports.Logger,
	productModule product.Module,
) (Server, error) {
	handler, err := router.NewRouter(cfg, logger, productModule.GetProductUseCase())
	if err != nil {
		return Server{}, err
	}

	return Server{
		logger: logger,
		Server: http.Server{
			Addr:              cfg.GetWebServerRestAddress(),
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}, nil
}

func (s *Server) Start(ctx context.Context) (err error) {
	go func() {
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(err.Error())
		}
	}()

	s.logger.Info(fmt.Sprintf("webserver available on: %v", s.Server.Addr))

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
