package base

import (
	"context"
	"strings"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.ProjectConf.AllowInit {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.PermissionDeny)
	}

	if l.svcCtx.Config.CoreRpc.Enabled {
		err = l.insertApiData()
		if err != nil {
			if strings.Contains(err.Error(), "common.createFailed") {
				return nil, errorx.NewInvalidArgumentError(i18n.AlreadyInit)
			}
			return nil, err
		}

		err = l.insertMenuData()
		if err != nil {
			return nil, err
		}

	}

	data, err := l.svcCtx.MmsRpc.InitDatabase(l.ctx, &mms.Empty{})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin Policy", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.DatabaseError)
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, data.Msg),
	}, nil
}
