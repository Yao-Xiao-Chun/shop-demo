package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	//搜索引擎配置
	MeiliSearch struct {
		Host      string
		MasterKey string
	}
	SearchConfig struct {
		Product string
	}
}
