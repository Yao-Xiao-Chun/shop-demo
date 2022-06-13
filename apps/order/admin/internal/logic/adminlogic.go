package logic

import (
	"context"

	"shop-demo/apps/order/admin/internal/svc"
	"shop-demo/apps/order/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.Request) (resp *types.Response, err error) {

	var res types.Response
	res.Message = "呵呵哒" + req.Name
	res.Code = 1
	data := make([]string, 0)
	data = append(data, "电网之外", "千里之旅")
	res.Data = data
	return &res, nil
}
