package memberrank

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-member-api/internal/logic/memberrank"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
)

// swagger:route post /member_rank/list memberrank GetMemberRankList
//
// Get member rank list | 获取会员等级列表
//
// Get member rank list | 获取会员等级列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MemberRankListReq
//
// Responses:
//  200: MemberRankListResp

func GetMemberRankListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberRankListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := memberrank.NewGetMemberRankListLogic(r.Context(), svcCtx)
		resp, err := l.GetMemberRankList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
