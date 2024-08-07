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

func getInt64PointerValue(ptr *int64) int64 {
	if ptr != nil {
		return *ptr
	}
	return int64(0)
}

func getStringPointerValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func getBoolPointerValue(ptr *bool) bool {
	if ptr != nil {
		return *ptr
	}
	return false
}

func (l *QueryUserDeviceListLogic) QueryUserDeviceList(req *types.DeviceListQueryReq) (resp *types.DeviceListResp, err error) {
	// todo: add your logic here and delete this line
	if req.UserId == "" {
		req.UserId = req.AddxUserId
	}

	member, err := l.svcCtx.MmsRpc.GetMemberById(l.ctx, &mms.UUIDReq{Id: req.UserId})

	if err != nil {
		l.Logger.Errorw("MmsRpc.GetMemberById error", logx.LogField{Key: "error", Value: err.Error()}, logx.LogField{Key: "req", Value: req})
		return &types.DeviceListResp{
			BaseDataInfo: types.BaseDataInfo{Code: -1, Msg: err.Error()},
			Data:         types.DeviceListInfo{},
		}, err
	}

	list, err := l.svcCtx.AddxProxy.QueryUserDeviceList(l.ctx, &synergy_addx_proxy_client.DeviceListQueryRequest{
		AddxUserId: "peckperk-" + member.GetForeinId(),
	})

	if err != nil {
		l.Logger.Errorw("AddxProxy.QueryUserDeviceList error", logx.LogField{Key: "error", Value: err.Error()}, logx.LogField{Key: "req", Value: req})
		return &types.DeviceListResp{
			BaseDataInfo: types.BaseDataInfo{
				Code: -1,
				Msg:  err.Error(),
			},

			Data: types.DeviceListInfo{
				BaseListInfo: types.BaseListInfo{
					Total: 0,
				},
				Data: []types.DeviceSummary{},
			},
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
			DeviceNetType:  getInt64PointerValue(summary.DeviceNetType),
			DeviceStatus:   getInt64PointerValue(summary.DeviceStatus),
			DeviceVipLevel: getInt64PointerValue(summary.DeviceVipLevel),
			FirmwareId:     getStringPointerValue(summary.FirmwareId),
			FirmwareStatus: getInt64PointerValue(summary.FirmwareStatus),
			Icon:           getStringPointerValue(summary.Icon),
			MacAddress:     getStringPointerValue(summary.MacAddress),
			Online:         getInt64PointerValue(summary.Online),
			PersonDetect:   getInt64PointerValue(summary.PersonDetect),
			SupportBirdVip: getBoolPointerValue(summary.SupportBirdVip),
		})
	}

	return &types.DeviceListResp{
		BaseDataInfo: types.BaseDataInfo{Code: int(list.Code), Msg: list.Message},
		Data: types.DeviceListInfo{
			BaseListInfo: types.BaseListInfo{
				Total: uint64(list.Count),
			},

			Data: summaryList,
		},
	}, err
}
