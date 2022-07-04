// Code generated by goctl. DO NOT EDIT!
// Source: seckill.proto

package server

import (
	"context"

	"shop-demo/apps/seckill/rpc/internal/logic"
	"shop-demo/apps/seckill/rpc/internal/svc"
	"shop-demo/apps/seckill/rpc/seckill"
)

type SecKillServer struct {
	svcCtx *svc.ServiceContext
	seckill.UnimplementedSecKillServer
}

func NewSecKillServer(svcCtx *svc.ServiceContext) *SecKillServer {
	return &SecKillServer{
		svcCtx: svcCtx,
	}
}

func (s *SecKillServer) CreateSecKill(ctx context.Context, in *seckill.SecKillRequest) (*seckill.SecKillResponse, error) {
	l := logic.NewCreateSecKillLogic(ctx, s.svcCtx)
	return l.CreateSecKill(in)
}
