package member

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMemberLogic {
	return &CreateMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMemberLogic) CreateMember(req *types.MemberInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.CreateMember(l.ctx,
		&mms.MemberInfo{
			Id:       req.Id,
			Status:   req.Status,
			Username: req.Username,
			Password: req.Password,
			Nickname: req.Nickname,
			RankId:   req.RankId,
			Mobile:   req.Mobile,
			Email:    req.Email,
			Avatar:   req.Avatar,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
