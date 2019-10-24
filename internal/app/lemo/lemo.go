package lemo

import (
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/pinguo-yinhailin/lemo/internal/pkg/config"
	"github.com/pinguo-yinhailin/lemo/internal/pkg/version"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type Lemo struct {
	app    *cli.App
	config *viper.Viper
	web    *iris.Application
}

func New() *Lemo {
	lemo := new(Lemo)

	app := cli.NewApp()
	app.Name = "lemo"
	app.Usage = "lemo 项目 web 服务示例"
	app.Before = lemo.before
	app.Action = lemo.action
	app.Version = version.String()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "env",
			Usage:  "项目运行环境，如 prod, dev, qa",
			EnvVar: "LEMO_ENV",
			Value:  "prod",
		},
	}
	lemo.app = app

	return lemo
}

func (l *Lemo) Run() {
	if err := l.app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func (l *Lemo) before(ctx *cli.Context) error {
	configPath := "../../configs"

	env := ctx.String("env")
	conf, err := config.Load(configPath, env)
	if err != nil {
		panic(err)
	}
	conf.Set("env", env)
	l.config = conf

	return nil
}

func (l *Lemo) action(ctx *cli.Context) error {
	web := iris.New()
	web.Logger().SetLevel(l.config.GetString("logLevel"))
	l.web = web
	l.registerRoutes()

	return web.Run(iris.Addr(":8369"), iris.WithoutServerError(iris.ErrServerClosed))
}
