package svc

import (
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/rpc/internal/config"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:        c,
		UserAuthModel: model.NewUserAuthModel(sqlConn),
		UserModel:     model.NewUserModel(sqlConn),
	}
}
