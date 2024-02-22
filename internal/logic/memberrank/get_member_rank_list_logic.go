package memberrank

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankListLogic {
	return &GetMemberRankListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberRankListLogic) GetMemberRankList(req *types.MemberRankListReq) (resp *types.MemberRankListResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMemberRankList(l.ctx,
		&mms.MemberRankListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Name:        req.Name,
			Description: req.Description,
			Remark:      req.Remark,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MemberRankListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MemberRankInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:       l.svcCtx.Trans.Trans(l.ctx, *v.Name),
				Name:        v.Name,
				Description: v.Description,
				Remark:      v.Remark,
				Code:        v.Code,
			})
	}
	return resp, nil
}
