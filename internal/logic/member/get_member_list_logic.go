package member

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberListLogic {
	return &GetMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberListLogic) GetMemberList(req *types.MemberListReq) (resp *types.MemberListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMemberList(l.ctx,
		&mms.MemberListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Username: req.Username,
			Nickname: req.Nickname,
			Mobile:   req.Mobile,
			Email:    req.Email,
			RankId:   req.RankId,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MemberListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MemberInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:   v.Status,
				Username: v.Username,
				Nickname: v.Nickname,
				RankId:   v.RankId,
				Mobile:   v.Mobile,
				Email:    v.Email,
				Avatar:   v.Avatar,
			})
	}
	return resp, nil
}
