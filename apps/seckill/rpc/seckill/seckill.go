// Code generated by goctl. DO NOT EDIT!
// Source: seckill.proto

package seckill

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SecKill interface {
		CreateSecKill(ctx context.Context, in *SecKillRequest, opts ...grpc.CallOption) (*SecKillResponse, error)
	}

	defaultSecKill struct {
		cli zrpc.Client
	}
)

func NewSecKill(cli zrpc.Client) SecKill {
	return &defaultSecKill{
		cli: cli,
	}
}

func (m *defaultSecKill) CreateSecKill(ctx context.Context, in *SecKillRequest, opts ...grpc.CallOption) (*SecKillResponse, error) {
	client := NewSecKillClient(m.cli.Conn())
	return client.CreateSecKill(ctx, in, opts...)
}