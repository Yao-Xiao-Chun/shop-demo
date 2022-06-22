package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OrderRPC   zrpc.RpcClientConf
	ProductRPC zrpc.RpcClientConf
	LogConf    logx.LogConf
	LoginRPC   zrpc.RpcClientConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}
