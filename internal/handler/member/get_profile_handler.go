package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/member"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
)

// swagger:route get /member/profile member GetProfile
//
// get users' own profile | 用户修改信息接口
//
// get users' own profile | 用户修改信息接口
//
// Responses:
//  200: LoginResp

func GetProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := member.NewGetProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetProfile()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
