package app

import (
	"context"
	"github.com/ljinf/template_project_v2/pkg/log"
	"github.com/ljinf/template_project_v2/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	name    string
	servers []server.Server
}

type Option func(a *App)

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = servers
	}
}

func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func (a *App) Run(ctx context.Context) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for _, srv := range a.servers {
		go func(srv server.Server) {
			err := srv.Start(ctx)
			if err != nil {
				log.Info(context.Background(), "Server start err", err.Error())
			}
		}(srv)
	}

	select {
	case <-signals:
		// Received termination signal
		log.Info(context.Background(), "Received termination signal")
	case <-ctx.Done():
		// Context canceled
		log.Info(context.Background(), "Context canceled")
	}

	// Gracefully stop the servers
	for _, srv := range a.servers {
		err := srv.Stop(ctx)
		if err != nil {
			log.Error(context.Background(), "Server stop err", err.Error())
		}
	}

	return nil
}
