package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"shop-demo/apps/login/rpc/login/login"
	"time"

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
	now := time.Now().Unix()

	//生成token
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, doLogin.UserId)

	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		Code:    int(doLogin.Code),
		Message: doLogin.Message,
		Token:   token,
	}, nil
}

func (l *LoginApiLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
