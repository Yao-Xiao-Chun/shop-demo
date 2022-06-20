package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"shop-demo/apps/product/rpc/product/internal/config"
	"shop-demo/apps/product/rpc/product/internal/server"
	"shop-demo/apps/product/rpc/product/internal/svc"
	"shop-demo/apps/product/rpc/product/product"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewProductServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		product.RegisterProductServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	fmt.Println(c.Mode)
	logx.DisableStat() //可以用来关闭 日志记录
	//logx.SetLevel(logx.ErrorLevel)// 日志记录等级

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
