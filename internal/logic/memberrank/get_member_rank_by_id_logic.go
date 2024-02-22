package memberrank

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberRankByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankByIdLogic {
	return &GetMemberRankByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberRankByIdLogic) GetMemberRankById(req *types.IDReq) (resp *types.MemberRankInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMemberRankById(l.ctx, &mms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MemberRankInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.MemberRankInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:        data.Name,
			Description: data.Description,
			Remark:      data.Remark,
			Code:        data.Code,
		},
	}, nil
}
