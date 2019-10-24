package lemo

import (
	"github.com/kataras/iris"
)

func (l *Lemo) beforeRoute(ctx iris.Context) {
	ctx.HTML("<p>before route</p>")

	ctx.Next()
}

func (l *Lemo) registerRoutes() {
	global := l.web.Party("/", l.beforeRoute)

	global.Get("/", l.Home)
	global.Get("/setting", l.Setting)

	article := global.Party("/articles")
	article.Get("/", l.Articles)
	article.Get("/{id:string}", l.Article)
}
