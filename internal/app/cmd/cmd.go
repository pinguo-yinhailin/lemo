package cmd

import (
	"log"
	"os"

	"github.com/pinguo-yinhailin/lemo/internal/pkg/config"
	"github.com/pinguo-yinhailin/lemo/internal/pkg/version"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type Commander interface {
	Command() cli.Command
}

type Cmd struct {
	app    *cli.App
	config *viper.Viper
}

func New() *Cmd {
	cmd := new(Cmd)

	app := cli.NewApp()
	app.Name = "lemo"
	app.Usage = "lemo 项目 cli 示例"
	app.Version = version.String()
	app.Before = cmd.before
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "env",
			Usage:  "项目运行环境，如 prod, dev, qa",
			EnvVar: "LEMO_ENV",
			Value:  "prod",
		},
	}
	cmd.app = app

	cmd.registerCommands()

	return cmd
}

func (c *Cmd) Run() {
	if err := c.app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func (c *Cmd) before(ctx *cli.Context) error {
	configPath := "../../configs"

	env := ctx.String("env")
	conf, err := config.Load(configPath, env)
	if err != nil {
		panic(err)
	}
	conf.Set("env", env)
	c.config = conf

	return nil
}
