package svc

import (
	"github.com/caichuanwang/go-zero-looklook/app/payment/cmd/rpc/internal/config"
	"github.com/caichuanwang/go-zero-looklook/app/payment/model"
	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config config.Config

	ThirdPaymentModel                  model.ThirdPaymentModel
	KqueuePaymentUpdatePayStatusClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
