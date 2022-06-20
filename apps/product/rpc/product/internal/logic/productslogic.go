package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"shop-demo/apps/product/model"
	"strconv"
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

// Products 获取商品列表
func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {

	products := make([]*product.ProductItem, 0) //不支持使用map int64的骚操作

	pdis := strings.Split(in.ProductIds, ",")

	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {

		v := item.(string)
		atoi, _ := strconv.Atoi(v)

		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, int64(atoi))
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

		logx.ErrorStack(err)

		return nil, err
	}

	for _, p := range ps.([]*model.Product) {
		products = append(products, &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		})
	}
	return &product.ProductResponse{Products: products}, nil
}
