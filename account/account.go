package main

import (
	"account_bend/account/internal/config"
	"account_bend/account/internal/handler"
	"account_bend/account/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/stat"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "account/etc/account-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	c.RestConf.Middlewares.Recover = false
	c.RestConf.Middlewares.Shedding = false
	c.RestConf.Middlewares.Log = true

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	go func() {
		server.Start()
	}()
	stat.DisableLog() // 关掉日志
	select {}
}
