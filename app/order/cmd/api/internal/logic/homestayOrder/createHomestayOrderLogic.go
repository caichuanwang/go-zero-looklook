package homestayOrder

import (
	"context"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/api/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/api/internal/types"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/rpc/order"
	"github.com/caichuanwang/go-zero-looklook/app/travel/cmd/rpc/pb"
	"github.com/caichuanwang/go-zero-looklook/common/ctxdata"
	"github.com/caichuanwang/go-zero-looklook/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateHomestayOrderLogic {
	return CreateHomestayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 创建订单
func (l *CreateHomestayOrderLogic) CreateHomestayOrder(req types.CreateHomestayOrderReq) (*types.CreateHomestayOrderResp, error) {
	homestayDetailRes, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &pb.HomestayDetailReq{Id: req.HomestayId})
	if err != nil {
		return nil, err
	}
	if homestayDetailRes == nil || homestayDetailRes.Homestay.Id == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("homestay no exists"), "CreateHomestayOrder homestay no exists id : %d", req.HomestayId)
	}

	userId := ctxdata.GetUidFromCtx(l.ctx)

	resp, err := l.svcCtx.OrderRpc.CreateHomestayOrder(l.ctx, &order.CreateHomestayOrderReq{
		HomestayId:    req.HomestayId,
		IsFood:        req.IsFood,
		LiveStartTime: req.LiveStartTime,
		LiveEndTime:   req.LiveEndTime,
		UserId:        userId,
		LivePeopleNum: req.LivePeopleNum,
		Remark:        req.Remark,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("create homestay order fail"), "create homestay order rpc CreateHomestayOrder fail req: %+v , err : %v ", req, err)
	}

	return &types.CreateHomestayOrderResp{
		OrderSn: resp.Sn,
	}, nil
}
