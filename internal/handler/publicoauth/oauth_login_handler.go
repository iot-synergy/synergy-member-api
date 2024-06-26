package publicoauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/publicoauth"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
)

// swagger:route post /oauth/login publicoauth OauthLogin
//
// Oauth log in | Oauth 登录
//
// Oauth log in | Oauth 登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OauthLoginReq
//
// Responses:
//  200: RedirectResp

func OauthLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthLoginReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicoauth.NewOauthLoginLogic(r.Context(), svcCtx)
		resp, err := l.OauthLogin(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
