package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-book/service/user/api/internal/config"
	"zero-book/service/user/model"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn,c.CacheRedis),
	}
}
