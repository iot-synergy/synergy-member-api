package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/device"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
)

// swagger:route post /device/updateVipActiveCode device UpdateVipActiveCode
//

//

//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: VipActiveCode
//
// Responses:
//  200: VipActiveCodeResp

func UpdateVipActiveCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VipActiveCode
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewUpdateVipActiveCodeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateVipActiveCode(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
