package logic

import (
	"context"
	"errors"
	"fmt"
	"shop-demo/apps/login/model"

	"shop-demo/apps/login/rpc/login/internal/svc"
	"shop-demo/apps/login/rpc/login/login"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoLoginLogic {
	return &DoLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DoLoginLogic) DoLogin(in *login.LoginRequest) (*login.LoginResponse, error) {
	//测试代码是否覆盖
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, 1) //查询id

	if err == model.ErrNotFound {
		return nil, errors.New("用户名不存在")
	}

	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return &login.LoginResponse{
		Code:    1,
		Message: "获取成功",
		Token:   "asdfjhk1",
	}, nil
}
