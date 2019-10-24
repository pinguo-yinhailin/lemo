package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/pinguo-yinhailin/lemo/internal/pkg/article"
	"github.com/urfave/cli"
)

type Article struct {
	cmd     *Cmd
	article article.Article
}

func NewArticle(cmd *Cmd) *Article {
	a := &Article{cmd: cmd}
	a.article = article.New()
	return a
}

func (a *Article) Command() cli.Command {
	return cli.Command{
		Name:      "article",
		ShortName: "a",
		Usage:     "文章管理cli，无子命令时默认显示文章统计信息",
		UsageText: "article [command options] [arguments...]",
		Description: `这里是一段非常详尽的命令使用信息
   执行 cmd article -h 或 cmd help article 时会用到...`,
		Action:      a.action,
		Subcommands: a.subcommands(),
	}
}

func (a *Article) action(ctx *cli.Context) error {
	fmt.Println("【文章统计信息】当前文章数：1322323，待审核数：239202")
	return nil
}

func (a *Article) subcommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "sync",
			Usage:     "批量同步文章信息",
			UsageText: "sync [command options] username password",
			ArgsUsage: "args usage",
			Before:    a.syncBefore,
			Action:    a.sync,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "max",
					Usage: "单次执行脚本允许的最大数据同步量",
					Value: 10,
				},
			},
		},
		{
			Name:      "push",
			Usage:     "并发推送文章到用户设备",
			UsageText: "push [command options]",
			Action:    a.push,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "max",
					Usage: "单次执行脚本允许的最大推送数量",
					Value: 10,
				},
			},
		},
	}
}

func (a *Article) syncBefore(ctx *cli.Context) error {
	if ctx.Args().Get(0) == "" {
		return fmt.Errorf("username 参数必填")
	}
	if ctx.Args().Get(1) == "" {
		return fmt.Errorf("password 参数必填")
	}
	return nil
}

func (a *Article) sync(ctx *cli.Context) error {
	fmt.Fprintf(ctx.App.Writer, "模拟数据同步...\n")
	for i := 1; i <= ctx.Int("max"); i++ {
		fmt.Fprintf(ctx.App.Writer, "同步文章 %v\n", i)
		a.article.Sync()
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Fprintf(ctx.App.Writer, "数据同步完成\n")
	return nil
}

func (a *Article) push(ctx *cli.Context) error {
	stream := make(chan string)
	done, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer close(stream)

		for {
			select {
			case <-done.Done():
				return
			case msg := <-stream:
				fmt.Fprintf(ctx.App.Writer, "%s\n", msg)
			}
		}
	}()

	return a.article.Push(done, stream, ctx.Int("max"))
}
