package member

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *LogoutLogic) Logout() (resp *types.BaseMsgResp, err error) {
	userId := l.ctx.Value("userId").(string)

	result, err := l.svcCtx.MmsRpc.BlockUserAllToken(l.ctx, &mms.UUIDReq{Id: userId})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: result.Msg}, nil
}
