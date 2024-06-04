package device

import (
	"context"

	"github.com/iot-synergy/synergy-addx-proxy/types/synergyAddxProxy"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenVipActiveCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenVipActiveCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenVipActiveCodeLogic {
	return &GenVipActiveCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GenVipActiveCodeLogic) GenVipActiveCode(req *types.VipActiveCodeGenReq) (resp *types.VipActiveCodeResp, err error) {

	rpcReq := &synergyAddxProxy.GenActivationCodeReq{
		ActivationDurationDay: req.ActivationDurationDay,
		From:                  req.From,
	}

	code, err := l.svcCtx.AddxProxy.GenActivationCode(l.ctx, rpcReq)

	if err != nil {
		l.Logger.Errorw("AddxProxy.GenActivationCode error", logx.LogField{Key: "error", Value: err.Error()}, logx.LogField{Key: "req", Value: req})
		return &types.VipActiveCodeResp{
			BaseMsgResp: types.BaseMsgResp{
				Code: -1,
				Msg:  err.Error(),
			},
		}, nil
	}

	return &types.VipActiveCodeResp{
		BaseMsgResp: types.BaseMsgResp{Code: 0, Msg: "ok"},
		Data: types.VipActiveCode{
			Id:             code.Id,
			ActivationCode: code.ActivationCode,
			State:          code.State,
			From:           code.From,
			DurationDay:    code.ActivationDurationDay,
			ActivationTime: code.ActivationTime,
			CreateAt:       code.CreatedAt,
			UpdateAt:       code.UpdatedAt,
		},
	}, nil
}
