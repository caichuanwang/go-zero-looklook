package logic

import (
	"context"

	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/rpc/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/usercenter/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthyUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAuthyUserIdResp{}, nil
}
