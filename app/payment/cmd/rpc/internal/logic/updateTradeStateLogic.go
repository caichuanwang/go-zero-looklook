package logic

import (
	"context"
	"encoding/json"
	"github.com/caichuanwang/go-zero-looklook/app/payment/model"
	"github.com/caichuanwang/go-zero-looklook/common/kqueue"
	"github.com/caichuanwang/go-zero-looklook/common/xerr"
	"github.com/pkg/errors"
	"time"

	"github.com/caichuanwang/go-zero-looklook/app/payment/cmd/rpc/internal/svc"
	"github.com/caichuanwang/go-zero-looklook/app/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeStateLogic {
	return &UpdateTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新交易状态
func (l *UpdateTradeStateLogic) UpdateTradeState(in *pb.UpdateTradeStateReq) (*pb.UpdateTradeStateResp, error) {
	thirdPayment, err := l.svcCtx.ThirdPaymentModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateTradeState FindOneBySn db err , sn : %s , err : %+v", in.Sn, err)
	}

	if thirdPayment == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("third payment record no exists"), " sn : %s", in.Sn)
	}
	//判断状态
	if in.PayStatus == model.ThirdPaymentPayTradeStateSuccess || in.PayStatus == model.ThirdPaymentPayTradeStateFAIL {
		//Want to modify as payment success, failure scenarios
		if thirdPayment.PayStatus != model.ThirdPaymentPayTradeStateWait {
			return &pb.UpdateTradeStateResp{}, nil
		}

	} else if in.PayStatus == model.ThirdPaymentPayTradeStateRefund {
		//Want to change to refund success scenario

		if thirdPayment.PayStatus != model.ThirdPaymentPayTradeStateSuccess {
			return nil, errors.Wrapf(xerr.NewErrMsg("Only orders with successful payment can be refunded"), "Only orders with successful payment can be refunded in : %+v", in)
		}
	} else {
		return nil, errors.Wrapf(xerr.NewErrMsg("This status is not currently supported"), "Modify payment flow status is not supported  in : %+v", in)
	}

	//更新
	thirdPayment.TradeState = in.TradeState
	thirdPayment.TransactionId = in.TransactionId
	thirdPayment.TradeType = in.TradeType
	thirdPayment.TradeStateDesc = in.TradeStateDesc
	thirdPayment.PayStatus = in.PayStatus
	thirdPayment.PayTime = time.Unix(in.PayTime, 0)
	if err := l.svcCtx.ThirdPaymentModel.UpdateWithVersion(l.ctx, nil, thirdPayment); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " UpdateTradeState UpdateWithVersion db  err:%v ,thirdPayment : %+v , in : %+v", err, thirdPayment, in)
	}
	if err := l.pubKqPaySuccess(in.Sn, in.PayStatus); err != nil {
		logx.WithContext(l.ctx).Errorf("l.pubKqPaySuccess : %+v", err)
	}
	return &pb.UpdateTradeStateResp{}, nil
}

func (l *UpdateTradeStateLogic) pubKqPaySuccess(orderSn string, payStatus int64) error {

	m := kqueue.ThirdPaymentUpdatePayStatusNotifyMessage{
		OrderSn:   orderSn,
		PayStatus: payStatus,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("kq UpdateTradeStateLogic pushKqPaySuccess task marshal error "), "kq UpdateTradeStateLogic pushKqPaySuccess task marshal error  , v : %+v", m)
	}

	return l.svcCtx.KqueuePaymentUpdatePayStatusClient.Push(string(body))
}
