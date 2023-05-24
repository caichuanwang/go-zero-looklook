package user

import (
	"context"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/rpc/usercenter"
	"github.com/caichuanwang/go-zero-looklook/common/ctxdata"
	"github.com/jinzhu/copier"

	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/api/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *DetailLogic) Detail(req types.UserInfoReq) (*types.UserInfoResp, error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	var userInfo types.User
	_ = copier.Copy(&userInfo, userInfoResp.User)

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
