package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/device"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
)

// swagger:route post /device/updateDeviceVipSet device UpdateDeviceVipSet
//

//

//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DeviceVipSetReq
//
// Responses:
//  200: DeviceVipSetUpdateResp

func UpdateDeviceVipSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceVipSetReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewUpdateDeviceVipSetLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDeviceVipSet(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
