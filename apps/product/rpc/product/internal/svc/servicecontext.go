package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shop-demo/apps/product/model"
	"shop-demo/apps/product/rpc/product/internal/config"
)

// ServiceContext 增加model
type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn),
	}
}
