package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shop-demo/apps/login/model"
	"shop-demo/apps/login/rpc/login/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis), //如果使用了redis不可以传空
	}
}
