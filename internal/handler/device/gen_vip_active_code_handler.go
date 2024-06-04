package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/device"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
)

// swagger:route post /device/genVipActiveCode device GenVipActiveCode
//

//

//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: VipActiveCodeGenReq
//
// Responses:
//  200: VipActiveCodeResp

func GenVipActiveCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VipActiveCodeGenReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewGenVipActiveCodeLogic(r.Context(), svcCtx)
		resp, err := l.GenVipActiveCode(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
