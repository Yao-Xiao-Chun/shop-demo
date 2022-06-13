package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shop-demo/apps/app/api/internal/config"
	"shop-demo/apps/order/rpc/order/order"
	"shop-demo/apps/product/rpc/product/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
