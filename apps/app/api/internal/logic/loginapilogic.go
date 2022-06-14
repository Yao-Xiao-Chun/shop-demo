package logic

import (
	"context"
	"fmt"
	"shop-demo/apps/login/rpc/login/login"

	"shop-demo/apps/app/api/internal/svc"
	"shop-demo/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginApiLogic {
	return &LoginApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginApiLogic) LoginApi(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 面对疾风二
	//调用login rpc 返回数据模型
	doLogin, err := l.svcCtx.LoginRPC.DoLogin(l.ctx, &login.LoginRequest{Account: req.Account, Password: req.Password})
	if err != nil {
		return nil, err
	}
	fmt.Println(doLogin)
	return &types.LoginResponse{
		Code:    int(doLogin.Code),
		Message: doLogin.Message,
		Token:   doLogin.Token,
	}, nil
}
