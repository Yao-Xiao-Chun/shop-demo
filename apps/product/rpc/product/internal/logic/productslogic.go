package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"strings"

	"shop-demo/apps/product/rpc/product/internal/svc"
	"shop-demo/apps/product/rpc/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {

	products := make(map[int64]*product.ProductItem)
	pdis := strings.Split(in.ProductIds, ",")
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		pid := item.(int64)
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, pid)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	for _, p := range ps.([]*model.Product) {
		products[p.Id] = &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		}
	}
	return &product.ProductResponse{Products: products}, nil
}
