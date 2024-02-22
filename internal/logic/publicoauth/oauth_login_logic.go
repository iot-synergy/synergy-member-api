package publicoauth

import (
	"context"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"
	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.RedirectResp, err error) {
	result, err := l.svcCtx.MmsRpc.OauthLogin(l.ctx, &mms.OauthLoginReq{
		State:    req.State,
		Provider: req.Provider,
	})
	if err != nil {
		return nil, err
	}

	return &types.RedirectResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data:         types.RedirectInfo{URL: result.Url},
	}, nil
}
