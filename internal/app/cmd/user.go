package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/pinguo-yinhailin/lemo/internal/pkg/user"
	"github.com/urfave/cli"
)

type User struct {
	cmd  *Cmd
	user user.User
}

func NewUser(cmd *Cmd) *User {
	u := &User{
		cmd:  cmd,
		user: user.New(),
	}

	return u
}

func (u *User) Command() cli.Command {
	return cli.Command{
		Name:      "user",
		ShortName: "u",
		Usage:     "虚拟用户管理cli",
		UsageText: "user [command options]",
		Description: `这里是一段非常详尽的命令使用信息
   执行 cmd user -h 或 cmd help user 时会用到...`,
		Action: u.action,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "quantity, q",
				Usage: "单次执行脚本生成的用户数量",
				Value: 5,
			},
		},
	}
}

func (u *User) action(ctx *cli.Context) error {
	done, cancel := context.WithCancel(context.Background())
	defer cancel()
	qty := ctx.Int("quantity")

	u.user.Run(done, u.cmd.config)
	go func() {
		for {
			select {
			case <-done.Done():
				return
			case p := <-u.user.C():
				fmt.Fprintf(u.cmd.app.Writer, "[3]生成用户：%s\n", p.Name)
			}
		}
	}()

	for {
		fmt.Fprintf(u.cmd.app.Writer, "开始生成虚拟用户，数量：%v >>>>>>>>>>>>>>>>>\n", qty)
		if err := u.user.Create(done, qty); err != nil {
			return err
		}
		time.Sleep(time.Second * 300)
	}
}
