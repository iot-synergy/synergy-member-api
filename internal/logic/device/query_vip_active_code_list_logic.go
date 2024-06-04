package device

import (
	"context"

	"github.com/iot-synergy/synergy-addx-proxy/types/synergyAddxProxy"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryVipActiveCodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryVipActiveCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryVipActiveCodeListLogic {
	return &QueryVipActiveCodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *QueryVipActiveCodeListLogic) QueryVipActiveCodeList(req *types.VipActiveCodeQueryReq) (resp *types.VipActiveCodeListResp, err error) {

	rpcReq := &synergyAddxProxy.QueryActivationCodeReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		State:    req.State,
		Sn:       req.SerialNumber,
	}

	dataList, err := l.svcCtx.AddxProxy.QueryActivationCode(l.ctx, rpcReq)

	if err != nil {
		l.Logger.Errorw("AddxProxy.QueryActivationCode error", logx.LogField{Key: "error", Value: err.Error()}, logx.LogField{Key: "req", Value: req})
		return nil, err
	}

	resp = &types.VipActiveCodeListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Total = uint64(dataList.Data.Count)

	for _, v := range dataList.Data.List {
		resp.Data = append(resp.Data,
			types.VipActiveCode{
				Id:             v.Id,
				ActivationCode: v.ActivationCode,
				State:          v.State,
				From:           v.From,
				SerialNumber:   v.Sn,
				OrderId:        v.OrderId,
				DurationDay:    v.ActivationDurationDay,
				ActivationTime: v.ActivationTime,
				CreateAt:       v.CreatedAt,
				UpdateAt:       v.UpdatedAt,
			})
	}

	return
}
