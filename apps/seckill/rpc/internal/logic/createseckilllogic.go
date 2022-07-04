package logic

import (
	"context"

	"shop-demo/apps/seckill/rpc/internal/svc"
	"shop-demo/apps/seckill/rpc/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSecKillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSecKillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSecKillLogic {
	return &CreateSecKillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSecKillLogic) CreateSecKill(in *seckill.SecKillRequest) (*seckill.SecKillResponse, error) {
	// todo: add your logic here and delete this line

	return &seckill.SecKillResponse{}, nil
}
