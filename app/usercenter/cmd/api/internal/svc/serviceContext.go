package svc

import (
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
