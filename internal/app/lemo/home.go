package lemo

import (
	"fmt"

	"github.com/kataras/iris"
)

func (l *Lemo) Home(ctx iris.Context) {
	ctx.HTML(fmt.Sprintf(`
<h1>current environment is %s</h1>
<p>GET <a href="/setting">/setting 显示配置文件信息</a></p>
<p>GET <a href="/articles">/articles 文章列表</a></p>
`, l.config.GetString("env")))
}
