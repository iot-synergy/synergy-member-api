package device

import (
	"context"

	"github.com/iot-synergy/synergy-addx-proxy/types/synergyAddxProxy"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLocalDevicesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLocalDevicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLocalDevicesLogic {
	return &QueryLocalDevicesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLocalDevicesLogic) QueryLocalDevices(req *types.QueryReq) (resp *types.DeviceListResp, err error) {
	// todo: add your logic here and delete this line

	re, er := l.svcCtx.AddxProxy.QueryLocalDevices(l.ctx, &synergyAddxProxy.LocalDevicesReq{
		Page:          req.Page,
		PageSize:      req.PageSize,
		Owner:         &req.Owner,
		ActivatedTime: &req.ActivatedTime,
	})

	if er != nil {
		return nil, er
	}

	resp = &types.DeviceListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Total = uint64(re.Count)

	for _, v := range re.Data.List {
		resp.Data = append(resp.Data,
			types.DeviceSummary{
				AddxId:         v.AddxUserId,
				SerialNumber:   v.SerialNumber,
				Activated:      v.Activated,
				ActivatedTime:  v.ActivatedTime,
				AdminName:      v.AdminName,
				DeviceName:     v.DeviceName,
				DeviceNetType:  v.GetDeviceNetType(),
				DeviceStatus:   v.GetDeviceStatus(),
				DeviceVipLevel: v.GetDeviceVipLevel(),
				FirmwareId:     v.GetFirmwareId(),
				FirmwareStatus: v.GetFirmwareStatus(),
				Icon:           v.GetIcon(),
				MacAddress:     v.GetMacAddress(),
				Online:         v.GetOnline(),
				PersonDetect:   v.GetPersonDetect(),
				SupportBirdVip: v.GetSupportBirdVip(),
			})
	}
	return resp, nil

}
