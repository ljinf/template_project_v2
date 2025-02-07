package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ljinf/template_project_v2/cmd/server/wire"
	"github.com/ljinf/template_project_v2/pkg/config"
	"github.com/ljinf/template_project_v2/pkg/log"
)

func main() {

	var envConf = flag.String("conf", "config/dev.yml", "config path, eg: -conf ./config/dev.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)
	log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	log.Info(context.Background(), fmt.Sprintf("server start host http://%s:%d",
		conf.GetString("app.http.host"), conf.GetInt("app.http.port")))
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}

}
