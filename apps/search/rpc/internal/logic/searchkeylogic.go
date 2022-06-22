package logic

import (
	"context"

	"shop-demo/apps/search/rpc/internal/svc"
	"shop-demo/apps/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchKeyLogic {
	return &SearchKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchKeyLogic) SearchKey(in *search.ProductSearchRequest) (*search.ProductSearchResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.SearchServe.Index("product_admin_list").AddDocuments("")

	if err != nil {
		return nil, err
	}
	return &search.ProductSearchResponse{}, nil
}
