//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ljinf/template_project_v2/internal/logic/appservice"
	"github.com/ljinf/template_project_v2/internal/logic/domainservice"
	"github.com/ljinf/template_project_v2/internal/logic/handler"
	"github.com/ljinf/template_project_v2/internal/repository"
	"github.com/ljinf/template_project_v2/internal/server"
	"github.com/ljinf/template_project_v2/pkg/app"
	"github.com/ljinf/template_project_v2/pkg/server/http"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var domainServiceSet = wire.NewSet(
	domainservice.NewUserDomainService,
)

var appServiceSet = wire.NewSet(
	appservice.NewUserAppService,
)

var handlerSet = wire.NewSet(
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	//server.NewJob,
	//server.NewTask,
)

// build App
func newApp(httpServer *http.Server /*,job *server.Job*/) *app.App {
	return app.NewApp(
		app.WithServer(httpServer /*,job*/),
		app.WithName("api-server"),
	)
}

func NewWire(*viper.Viper) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		domainServiceSet,
		appServiceSet,
		handlerSet,
		serverSet,
		newApp,
	))
}
