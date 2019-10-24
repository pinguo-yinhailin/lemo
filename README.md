# Go 项目规范demo

文档地址：[http://gitlab.camera360.com/yinhailin/server-wiki/tree/master/项目规范/go](http://gitlab.camera360.com/yinhailin/server-wiki/tree/master/%E9%A1%B9%E7%9B%AE%E8%A7%84%E8%8C%83/go)

# demo 说明

该 demo 包含两个可执行程序，分别是：
 
 - `cmd/cmd` cli程序，模拟服务端定时脚本等
 - `cmd/lemo` 模拟http服务
 
 进入相应的应用目录，执行命令：
 
```bash
cd cmd/cmd
go build
./cmd -h
```

即可看到使用帮助信息