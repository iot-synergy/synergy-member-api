package device

import (
	"context"
	"github.com/iot-synergy/synergy-addx-proxy/synergy_addx_proxy_client"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeviceVipSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeviceVipSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeviceVipSetLogic {
	return &UpdateDeviceVipSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDeviceVipSetLogic) UpdateDeviceVipSet(req *types.DeviceVipSetReq) (resp *types.DeviceVipSetUpdateResp, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.AddxProxy.UpdateDeviceVipSet(l.ctx, &synergy_addx_proxy_client.DeviceVipSetRequest{
		SerialNumber: req.SerialNumber,
		ProductId:    req.ProductId,
	})

	if err != nil {
		return &types.DeviceVipSetUpdateResp{
			BaseMsgResp: types.BaseMsgResp{
				Code: -1,
				Msg:  err.Error(),
			},
			Data: "",
		}, err
	}

	return &types.DeviceVipSetUpdateResp{
		BaseMsgResp: types.BaseMsgResp{
			Code: int(response.Code),
			Msg:  response.Message,
		},
		Data: response.ProductId,
	}, nil
}
