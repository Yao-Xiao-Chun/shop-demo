package logic

import (
	"context"

	"shop-demo/apps/login/rpc/login/internal/svc"
	"shop-demo/apps/login/rpc/login/login"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginOutLogic {
	return &LoginOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginOutLogic) LoginOut(in *login.LoginRequest) (*login.LoginRequest, error) {
	// todo: add your logic here and delete this line

	return &login.LoginRequest{}, nil
}
