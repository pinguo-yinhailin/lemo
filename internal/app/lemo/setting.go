package lemo

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/pinguo-yinhailin/lemo/internal/pkg/model"
)

func (l *Lemo) Setting(ctx iris.Context) {
	conf := new(model.SettingConfig)
	l.config.UnmarshalKey("system.setting", conf)

	ctx.HTML(fmt.Sprintf(`
<h1>%s</h1>
<p><a href="/">首页</a></p>
<p>超级管理员：%v</p>
<p>最大分页：%v</p>
<p>圆周率：%v</p>
<p>联系邮箱：%+v</p>
`, conf.Name, conf.IsAdmin, conf.MaxPage, conf.Pi, conf.Emails))
}
