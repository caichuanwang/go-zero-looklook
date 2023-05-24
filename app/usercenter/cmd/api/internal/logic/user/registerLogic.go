package user

import (
	"context"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/rpc/usercenter"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/api/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterResp, error) {

	registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)

	return &resp, nil
}
