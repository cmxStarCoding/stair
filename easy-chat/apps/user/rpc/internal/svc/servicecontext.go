package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"stair/easy-chat/apps/user/models"
	"stair/easy-chat/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
