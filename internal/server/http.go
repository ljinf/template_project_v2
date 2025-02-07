package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project_v2/internal/logic/handler"
	"github.com/ljinf/template_project_v2/internal/middleware"
	"github.com/ljinf/template_project_v2/internal/router"
	"github.com/ljinf/template_project_v2/pkg/enum"
	"github.com/ljinf/template_project_v2/pkg/server/http"
	"github.com/spf13/viper"
)

func NewHTTPServer(
	conf *viper.Viper,
	userHandler handler.UserHandler,
) *http.Server {
	mode := gin.DebugMode
	if conf.GetString("app.env") == enum.ModeProd {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	s := http.NewServer(
		gin.New(),
		http.WithServerHost(conf.GetString("app.http.host")),
		http.WithServerPort(conf.GetInt("app.http.port")),
	)

	s.Use(
		middleware.CORSMiddleware(),
		middleware.StartTrace(),
		middleware.LogAccess(),
		middleware.GinPanicRecovery(),
	)

	v1 := s.Group("/v1")
	router.RegistryRouter(v1, userHandler)

	return s
}
