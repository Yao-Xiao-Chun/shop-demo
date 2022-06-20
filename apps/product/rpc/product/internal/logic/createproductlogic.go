package logic

import (
	"context"
	"shop-demo/apps/product/model"
	"shop-demo/apps/product/rpc/product/product"

	"shop-demo/apps/product/rpc/product/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.ProductDataRequest) (*product.ProductDataResponse, error) {
	//获取数据写入
	productModel := model.Product{
		Cateid:   1,
		Name:     in.Name,
		Subtitle: in.SubTitle,
		Images:   in.Images,
		Price:    float64(in.Price),
		Stock:    in.Stock,
		Status:   in.Status,
	}

	//创建商品
	insert, err := l.svcCtx.ProductModel.Insert(l.ctx, &productModel)

	if err != nil {
		return nil, err
	}

	id, _ := insert.LastInsertId()

	return &product.ProductDataResponse{Res: true, GoodsId: id}, nil
}
