package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project_v2/pkg/log"
	"net/http"
	"time"
)

type Server struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
}
type Option func(s *Server)

func NewServer(engine *gin.Engine, opts ...Option) *Server {
	s := &Server{
		Engine: engine,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}
func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s,
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(context.Background(), "listen", err)
	}

	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	log.Info(context.Background(), "Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		log.Fatal(context.Background(), "Server forced to shutdown: ", err)
	}

	log.Info(context.Background(), "Server exiting")
	return nil
}
