package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shop-demo/apps/product/model"
	"shop-demo/apps/product/rpc/product/internal/config"
)

// ServiceContext 增加model
type ServiceContext struct {
	Config        config.Config
	ProductModel  model.ProductModel  //商品
	CategoryModel model.CategoryModel //商品分类
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource) //数据库连接资源
	return &ServiceContext{
		Config:        c,
		ProductModel:  model.NewProductModel(conn),
		CategoryModel: model.NewCategoryModel(conn),
	}
}
