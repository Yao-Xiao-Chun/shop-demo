package common

import (
	"context"
	"sync"
)

// 消息
type msg struct {
	key string
	val interface{}
}

type Options struct {
	check  func()
	apply  func(val interface{})
	worker int
	buffer int
}

type Batcher struct {
	opts     Options
	Do       func(ctx context.Context, val map[string][]interface{}) //满足聚合条件就执行do方法，其中val 参数为聚合后的数据
	Sharding func(key string) int
	chans    []chan *msg
	wait     sync.WaitGroup
}

func New(opts ...Options) *Batcher {

	b := &Batcher{}
	for _, opt := range opts {
		opt.apply(&b.opts)
	}
	b.opts.check()
	b.chans = make([]chan *msg, b.opts.worker)

	for i := 0; i < b.opts.worker; i++ {
		b.chans[i] = make(chan *msg, b.opts.buffer)
	}

	return b
}
