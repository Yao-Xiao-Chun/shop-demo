package logic

import (
	"context"

	"shop-demo/apps/order/rpc/order/internal/svc"
	"shop-demo/apps/order/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	res := order.OrdersResponse{}
	infos := make([]*order.OrderItem, 0)
	tmp := order.OrderItem{
		OrderId:   "1",
		Quantity:  12,
		Payment:   12.3,
		ProductId: 0,
	}
	infos = append(infos, &tmp)
	res.Orders = infos
	res.IsEnd = false
	res.CreateTime = "2022-06-13"

	return &res, nil
}
