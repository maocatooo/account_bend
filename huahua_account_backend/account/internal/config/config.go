package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
	Mysql struct {
		DataSource string
	}
	Wechat WeChatConfig

	//CacheRedis cache.CacheConf
}

type WeChatConfig struct {
	AppID     string
	AppSecret string
}
