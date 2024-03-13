package member

import (
	"context"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncFirebaseMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncFirebaseMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncFirebaseMembersLogic {
	return &SyncFirebaseMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SyncFirebaseMembersLogic) SyncFirebaseMembers() (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.SyncFirebaseMember(l.ctx,
		&mms.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
