package lemo

import (
	"fmt"

	"github.com/kataras/iris"
)

func (l *Lemo) Articles(ctx iris.Context) {
	ctx.HTML(`
<p><a href="/">首页</a></p>
<p><a href="/articles/1">在 Go 中使用 Viper 加载配置文件(1)</a><p>
<p><a href="/articles/2">在 Go 中使用 Viper 加载配置文件(2)</a><p>
<p><a href="/articles/3">在 Go 中使用 Viper 加载配置文件(3)</a><p>
<p><a href="/articles/4">在 Go 中使用 Viper 加载配置文件(4)</a><p>
<p><a href="/articles/5">在 Go 中使用 Viper 加载配置文件(5)</a><p>
<p><a href="/articles/6">在 Go 中使用 Viper 加载配置文件(6)</a><p>
<p><a href="/articles/7">在 Go 中使用 Viper 加载配置文件(7)</a><p>
`)
}

func (l *Lemo) Article(ctx iris.Context) {
	ctx.HTML(fmt.Sprintf(`
<p><a href="/">首页</a></p>
<p><a href="/articles">列表</a></p>
<h1>在 Go 中使用 Viper 加载配置文件(%v)<h1>
<article>文章正文</article>
`, ctx.Params().Get("id")))
}
