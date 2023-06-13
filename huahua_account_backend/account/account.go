package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/stat"
	"huahua_account_backend/account/internal/config"
	"huahua_account_backend/account/internal/handler"
	"huahua_account_backend/account/internal/svc"

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
	c.RestConf.Middlewares.Log = false

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
