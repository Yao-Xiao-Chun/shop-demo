// Code generated by goctl. DO NOT EDIT!
// Source: login.proto

package login

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Login interface {
		// 登录
		DoLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		// 登出
		LoginOut(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginRequest, error)
	}

	defaultLogin struct {
		cli zrpc.Client
	}
)

func NewLogin(cli zrpc.Client) Login {
	return &defaultLogin{
		cli: cli,
	}
}

// 登录
func (m *defaultLogin) DoLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := NewLoginClient(m.cli.Conn())
	return client.DoLogin(ctx, in, opts...)
}

// 登出
func (m *defaultLogin) LoginOut(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginRequest, error) {
	client := NewLoginClient(m.cli.Conn())
	return client.LoginOut(ctx, in, opts...)
}
