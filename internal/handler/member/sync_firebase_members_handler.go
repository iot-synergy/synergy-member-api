package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/member"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
)

// swagger:route post /member/sync_firebase_members member SyncFirebaseMembers
//
// Sync firebase Member Data | 同步firebase上的用户数据
//
// Sync firebase Member Data | 同步firebase上的用户数据
//
// Responses:
//  200: BaseMsgResp

func SyncFirebaseMembersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := member.NewSyncFirebaseMembersLogic(r.Context(), svcCtx)
		resp, err := l.SyncFirebaseMembers()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
