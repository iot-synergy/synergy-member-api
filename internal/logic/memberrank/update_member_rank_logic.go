package memberrank

import (
	"context"

	"github.com/iot-synergy/synergy-member-api/internal/logic/publicmember"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRankLogic {
	return &UpdateMemberRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMemberRankLogic) UpdateMemberRank(req *types.MemberRankInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.UpdateMemberRank(l.ctx,
		&mms.MemberRankInfo{
			Id:          req.Id,
			Name:        req.Name,
			Description: req.Description,
			Remark:      req.Remark,
			Code:        req.Code,
		})
	if err != nil {
		return nil, err
	}

	publicmember.MemberRankData = make(map[uint64]string)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
