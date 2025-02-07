package grpc

import (
	"context"
	"fmt"
	"github.com/ljinf/template_project_v2/pkg/log"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	*grpc.Server
	host string
	port int
}

type Option func(s *Server)

func NewServer(opts ...Option) *Server {
	s := &Server{
		Server: grpc.NewServer(),
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
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Fatal(context.Background(), "Failed to listen:", err.Error())
	}
	if err = s.Server.Serve(lis); err != nil {
		log.Fatal(context.Background(), "Failed to serve:", err.Error())
	}
	return nil

}
func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	s.Server.GracefulStop()

	log.Info(context.Background(), "Server exiting")

	return nil
}
