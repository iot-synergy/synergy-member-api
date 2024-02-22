package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/member"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
)

// swagger:route get /member/logout member Logout
//
// Log out the service | 用户登出
//
// Log out the service | 用户登出
//
// Responses:
//  200: BaseMsgResp

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := member.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
