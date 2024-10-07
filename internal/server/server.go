package server

import (
	"context"
	"github.com/zenorachi/youtube-task/internal/config"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(cfg *config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    cfg.HTTPAddress(),
	}

	return &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *Server) Run() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown(ctx context.Context) error {
	c, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(c)
}
