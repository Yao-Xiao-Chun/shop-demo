package logic

import (
	"context"

	"shop-demo/apps/order/rpc/order/internal/svc"
	"shop-demo/apps/order/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	// todo: add your logic here and delete this line

	return &order.OrdersResponse{}, nil
}
