package device

import (
	"context"

	"github.com/iot-synergy/synergy-addx-proxy/types/synergyAddxProxy"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVipActiveCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVipActiveCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVipActiveCodeLogic {
	return &UpdateVipActiveCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateVipActiveCodeLogic) UpdateVipActiveCode(req *types.VipActiveCode) (resp *types.VipActiveCodeResp, err error) {
	rpcReq := &synergyAddxProxy.ActivationCodeInfo{
		Id:                    req.Id,
		ActivationCode:        req.ActivationCode,
		CreatedAt:             0,
		UpdatedAt:             0,
		State:                 req.State,
		From:                  req.From,
		Sn:                    req.SerialNumber,
		OrderId:               req.OrderId,
		Email:                 req.Email,
		ActivationTime:        req.ActivationTime,
		ActivationDurationDay: req.DurationDay,
	}

	code, err := l.svcCtx.AddxProxy.UpdateActivationCode(l.ctx, rpcReq)

	if err != nil {
		l.Logger.Errorw("AddxProxy.UpdateActivationCode error", logx.LogField{Key: "error", Value: err.Error()}, logx.LogField{Key: "req", Value: req})
		return nil, err
	}

	return &types.VipActiveCodeResp{
		BaseMsgResp: types.BaseMsgResp{
			Code: 0,
			Msg:  "OK",
		},
		Data: types.VipActiveCode{
			Id:             code.Id,
			ActivationCode: code.ActivationCode,
			State:          code.State,
			From:           code.From,
			Email:          code.Email,
			SerialNumber:   code.Sn,
			OrderId:        code.OrderId,
			DurationDay:    code.ActivationDurationDay,
			ActivationTime: code.ActivationTime,
			CreateAt:       code.CreatedAt,
			UpdateAt:       code.UpdatedAt,
		},
	}, nil
}
