package device

import (
	"context"
	"github.com/iot-synergy/synergy-addx-proxy/synergy_addx_proxy_client"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserDeviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDeviceListLogic {
	return &QueryUserDeviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserDeviceListLogic) QueryUserDeviceList(req *types.DeviceListQueryReq) (resp *types.DeviceListResp, err error) {
	// todo: add your logic here and delete this line
	member, err := l.svcCtx.MmsRpc.GetMemberById(l.ctx, &mms.UUIDReq{Id: req.UserId})
	list, err := l.svcCtx.AddxProxy.QueryUserDeviceList(l.ctx, &synergy_addx_proxy_client.DeviceListQueryRequest{
		AddxUserId:   "peckperk-" + member.GetForeinId(),
		SerialNumber: &req.SerialNumber,
	})
	if err != nil {
		return &types.DeviceListResp{
			BaseDataInfo: types.BaseDataInfo{
				Code: -1,
				Msg:  err.Error(),
			},
			BaseListInfo: types.BaseListInfo{
				Total: 0,
			},
			Data: nil,
		}, err
	}

	var summaryList []types.DeviceSummary
	for _, summary := range list.Data.List {
		summaryList = append(summaryList, types.DeviceSummary{
			AddxId:         summary.AddxUserId,
			SerialNumber:   summary.SerialNumber,
			Activated:      summary.Activated,
			ActivatedTime:  summary.ActivatedTime,
			AdminName:      summary.AdminName,
			DeviceName:     summary.DeviceName,
			DeviceNetType:  *summary.DeviceNetType,
			DeviceStatus:   *summary.DeviceStatus,
			DeviceVipLevel: *summary.DeviceVipLevel,
			FirmwareId:     *summary.FirmwareId,
			FirmwareStatus: *summary.FirmwareStatus,
			Icon:           *summary.Icon,
			MacAddress:     *summary.MacAddress,
			Online:         *summary.Online,
			PersonDetect:   *summary.PersonDetect,
			SupportBirdVip: *summary.SupportBirdVip,
		})
	}

	return &types.DeviceListResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: int(list.Code),
			Msg:  list.Message,
		},
		BaseListInfo: types.BaseListInfo{
			Total: uint64(list.Count),
		},
		Data: summaryList,
	}, err
}
