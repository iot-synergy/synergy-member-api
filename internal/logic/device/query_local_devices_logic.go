package device

import (
	"context"

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

func (l *QueryLocalDevicesLogic) QueryLocalDevices() (resp *types.DeviceListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
