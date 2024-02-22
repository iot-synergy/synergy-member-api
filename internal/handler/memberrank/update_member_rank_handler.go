package memberrank

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/memberrank"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
)

// swagger:route post /member_rank/update memberrank UpdateMemberRank
//
// Update member rank information | 更新会员等级
//
// Update member rank information | 更新会员等级
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MemberRankInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateMemberRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberRankInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := memberrank.NewUpdateMemberRankLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMemberRank(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
