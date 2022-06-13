package logic

import (
	"context"

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
	// todo: add your logic here and delete this line
	var res product.ProductResponse
	tmp := product.ProductItem{}
	tmp.ProductId = 1
	tmp.Name = "测试"
	tmp.Description = "描述哦"
	infos := make([]*product.ProductItem, 0)
	infos = append(infos, &tmp)
	res.Products = infos
	return &res, nil
}
