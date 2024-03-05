package device

import (
	"net/http"

	"github.com/iot-synergy/synergy-member-api/internal/logic/device"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryLocalDevicesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := device.NewQueryLocalDevicesLogic(r.Context(), svcCtx)
		resp, err := l.QueryLocalDevices()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
