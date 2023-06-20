package homestayOrder

import (
	"context"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/api/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/api/internal/types"
	"github.com/caichuanwang/go-zero-looklook/app/order/cmd/rpc/order"
	"github.com/caichuanwang/go-zero-looklook/common/ctxdata"
	"github.com/caichuanwang/go-zero-looklook/common/tool"
	"github.com/caichuanwang/go-zero-looklook/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserHomestayOrderListLogic {
	return UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l UserHomestayOrderListLogic) UserHomestayOrderList(req types.UserHomestayOrderListReq) (*types.UserHomestayOrderListResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &order.UserHomestayOrderListReq{
		UserId:      userId,
		TraderState: req.TradeState,
		PageSize:    req.PageSize,
		LastId:      req.LastId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get user homestay order list"), "Failed to get user homestay order list err : %v ,req:%+v", err, req)
	}

	var typesUserHomestayOrderList []types.UserHomestayOrderListView
	if len(resp.List) > 0 {

		for _, homestayOrder := range resp.List {

			var typeHomestayOrder types.UserHomestayOrderListView
			_ = copier.Copy(&typeHomestayOrder, homestayOrder)

			typeHomestayOrder.OrderTotalPrice = tool.Fen2Yuan(homestayOrder.OrderTotalPrice)

			typesUserHomestayOrderList = append(typesUserHomestayOrderList, typeHomestayOrder)
		}
	}

	return &types.UserHomestayOrderListResp{
		List: typesUserHomestayOrderList,
	}, nil
}
