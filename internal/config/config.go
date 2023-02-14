package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache cache.CacheConf
	Mysql struct {
		Datasource string
	}
	COS struct {
		BucketName string
		AppId      string
		Region     string
		SecretId   string
		SecretKey  string
	}
}
